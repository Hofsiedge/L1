package sleep

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall" // deprecated, but good enough for this task
	"time"
)

type sleeping struct {
	endTime time.Time
	channel chan struct{}
}
type Sleeper struct {
	sync.Mutex
	nextAlarm       time.Time
	blockedChannels []*sleeping
	signalChannel   chan os.Signal
	done            chan struct{}
	closed          chan struct{}
}

var maxTime time.Time = time.Date(20000, time.January, 1, 0, 0, 0, 0, time.UTC)

// monitor handles alarm signals and waking up
func (s *Sleeper) monitor() {
	// alarm calls cancel previous calls, so the logic is a bit contrived
	for {
		select {
		// on closing
		case <-s.done:
			s.Lock()
			defer s.Unlock()
			// waking up all the sleeping goroutines
			for _, sleeping := range s.blockedChannels {
				sleeping.channel <- struct{}{}
			}
			close(s.closed)
			return
		// on alarm/sleeping event
		case <-s.signalChannel:
			// locking to prevent data races
			s.Lock()
			currentTime := time.Now()
			blocked := make([]*sleeping, 0, len(s.blockedChannels))
			s.nextAlarm = maxTime

			// check every sleeping
			for _, sleeping := range s.blockedChannels {
				// if this one is meant to be asleep
				if sleeping.endTime.After(currentTime) {
					// leave it blocked
					blocked = append(blocked, sleeping)
					// potentially update time to next alarm
					if sleeping.endTime.Before(s.nextAlarm) {
						s.nextAlarm = sleeping.endTime
					}
				} else {
					// sending a signal to wake up
					sleeping.channel <- struct{}{}
				}
			}
			s.blockedChannels = blocked
			// calling alarm again
			// alarm signal will be received after the closest waking up time
			secondsLeft := s.nextAlarm.Sub(currentTime).Seconds()
			secondsUntilNextAlarm := uintptr(math.Ceil(float64(secondsLeft)))
			_, _, err := syscall.Syscall(
				syscall.SYS_ALARM, secondsUntilNextAlarm, 0, 0)
			if err != syscall.Errno(0) {
				log.Printf("error in the sleep function: %v", err)
			}

			s.Unlock()
		}
	}
}

// should only be called once.
//
// calling it multiple times might result in leaks.
func NewSleeper() *Sleeper {
	sleeper := Sleeper{
		signalChannel:   make(chan os.Signal, 1),
		blockedChannels: make([]*sleeping, 0),
		done:            make(chan struct{}),
		closed:          make(chan struct{}),
		nextAlarm:       maxTime,
	}
	signal.Notify(sleeper.signalChannel, syscall.SIGALRM)

	go sleeper.monitor()
	return &sleeper
}

// Close wakes up all the sleeping goroutines and closes Sleeper
func (s *Sleeper) Close() {
	close(s.done)
	<-s.closed
}

// Sleep is not too accurate but should be able
// to provide 1 second precision.
//
// Only accepts integer seconds since so does the underlying mechanism
// (alarm syscall)
func (s *Sleeper) Sleep(seconds int) {
	if seconds <= 0 {
		return
	}
	// locking to ensure proper data handling
	s.Lock()
	endTime := time.Now().Add(time.Second * time.Duration(seconds))
	channel := make(chan struct{})
	s.blockedChannels = append(s.blockedChannels, &sleeping{endTime, channel})
	nextAlarm := s.nextAlarm
	s.Unlock()
	fmt.Printf("unlocked, %v\n", nextAlarm)
	if endTime.Before(nextAlarm) {
		// pushing an alarm signal to reissue a new alarm syscall
		s.signalChannel <- syscall.SIGALRM
	}
	<-channel
}

// naive implementation. all sleeping goroutines wake up at the same time
// due to alarm syscall cancellation on new syscall
/*
// SigalrmSleep uses `alarm` syscall.
func SigalrmSleep(seconds int) {
	if seconds < 0 {
		seconds = 0
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGALRM)
	arg := uint(seconds)
	result, r, err := syscall.Syscall(syscall.SYS_ALARM, uintptr(arg), 0, 0)
	fmt.Println(result, r, err)
	<-sig
}
*/
