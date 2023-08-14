package main

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования)
*/

import "fmt"

type Human struct {
	firstName string
	lastName  string
}

// concatenate first and last names
func (h Human) GetFullName() string {
	return h.firstName + " " + h.lastName
}

// Action contains a name of the action and its executor
type Action struct {
	Human
	actionName string
}

func (a Action) Announce() string {
	return fmt.Sprintf("%s is about to %s", a.GetFullName(), a.actionName)
}

func main() {
	elonMusk := Human{
		firstName: "Elon",
		lastName:  "Musk",
	}
	elonMuskTweets := Action{
		Human:      elonMusk,
		actionName: "make a tweet",
	}
	fmt.Println(elonMuskTweets.Announce())
}
