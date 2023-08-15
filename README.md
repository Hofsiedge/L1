# L1

## Tasks
| №  | main                        | library                                               | test                                                                                                                                                                                                                                                                         | task |
|----|-----------------------------|-------------------------------------------------------|----------------------------------------------|--------------------------                                                                                                                                                                                                     |
| 1  | [main](cmd/task_1/main.go ) |                                                       |                                              | Дана структура Human. Реализовать встраивание методов в структуре Action от родительской структуры Human                                                                                                                      |
| 2  | [main](cmd/task_2/main.go ) |                                                       |                                              | Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива и выведет их квадраты в stdout                                                                                                  |
| 3  | [main](cmd/task_3/main.go ) |                                                       |                                              | Дана последовательность чисел. Найти сумму их квадратов с использованием конкурентных вычислений                                                                                                                              |
| 4  | [main](cmd/task_4/main.go ) |                                                       |                                              | Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте. |
| 5  | [main](cmd/task_5/main.go ) |                                                       |                                              | Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться                                                      |
| 6  | [main](cmd/task_6/main.go ) |                                                       |                                              | Реализовать все возможные способы остановки выполнения горутины                                                                                                                                                               |
| 7  | [main](cmd/task_7/main.go ) |                                                       |                                              | Реализовать конкурентную запись данных в map                                                                                                                                                                                  |
| 8  | [main](cmd/task_8/main.go ) |                                                       | [test](cmd/task_8/setbit_test.go)            | Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0                                                                                                                                          |
| 9  | [main](cmd/task_9/main.go ) | [internal/pipeline](internal/pipeline/pipeline.go)    |                                              | Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout                                        |
| 10 | [main](cmd/task_10/main.go) |                                                       |                                              | Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна                    |
| 11 | [main](cmd/task_11/main.go) |                                                       |                                              | Реализовать пересечение двух неупорядоченных множеств                                                                                                                                                                         |
| 12 | [main](cmd/task_12/main.go) |                                                       |                                              | Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество                                                                                                                           |
| 13 | [main](cmd/task_13/main.go) |                                                       |                                              | Поменять местами два числа без создания временной переменной                                                                                                                                                                  |
| 14 | [main](cmd/task_14/main.go) |                                                       |                                              | Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}                                                                                       |
| 15 | [main](cmd/task_15/main.go) |                                                       |                                              | К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.                                                                                             |
| 16 | [main](cmd/task_16/main.go) | [internal/quicksort](internal/quicksort/quicksort.go) | [test](internal/quicksort/quicksort_test.go) | Реализовать быструю сортировку массива (quicksort) встроенными методами языка                                                                                                                                                 |
| 17 | [main](cmd/task_17/main.go) | [internal/bisect](internal/bisect/bisect.go)          | [test](internal/bisect/bisect_test.go)       | Реализовать бинарный поиск встроенными методами языка                                                                                                                                                                         |
| 18 | [main](cmd/task_18/main.go) |                                                       |                                              | Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика                                                                      |
| 19 | [main](cmd/task_19/main.go) |                                                       |                                              | Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.                                                                                         |
| 20 | [main](cmd/task_20/main.go) |                                                       |                                              | Разработать программу, которая переворачивает слова в строке.  Пример: «snow dog sun — sun dog snow».                                                                                                                         |
| 21 | [main](cmd/task_21/main.go) |                                                       |                                              | Реализовать паттерн «адаптер» на любом примере                                                                                                                                                                                |
| 22 | [main](cmd/task_22/main.go) |                                                       |                                              | Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20                                                                                                  |
| 23 | [main](cmd/task_23/main.go) |                                                       |                                              | Удалить i-ый элемент из слайса                                                                                                                                                                                                |
| 24 | [main](cmd/task_24/main.go) |                                                       |                                              | Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.                                                            |
| 25 | [main](cmd/task_25/main.go) | [internal/sleep](internal/sleep/sleep.go)             |                                              | Реализовать собственную функцию sleep                                                                                                                                                                                         |
| 26 | [main](cmd/task_26/main.go) |                                                       |                                              | Разработать программу, которая проверяет, что все символы в строке уникальные. Функция проверки должна быть регистронезависимой                                                                                               |

## Colloquial questions

### Какой самый эффективный способ конкатенации строк?
According to [docs](https://pkg.go.dev/strings#Builder), `strings.Builder` is optimized for for efficient `string` building.

Usage example:
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder

	fragments := []string{"fragment 1", "fragment 2", "fragment 3"}
	for _, fragment := range fragments {
		b.WriteString(fragment)
		b.WriteString(". ")
	}
	fmt.Println(b.String()) // fragment 1. fragment 2. fragment 3.
}
```

### Что такое интерфейсы, как они применяются в Go?
Interfaces are type contracts that declare what methods must be implemented.
Interfaces are implemented implicitly in Go.
There are 3 different ways of using interfaces.

The first one is to use an empty interface `interface{}` that is implemented
by everything in Go to make a function accept anything as an argument.
`interface{}` is also aliased as `any` and usually it is more convenient
to use this form.

The second one is to use an interface to specify the methods that the type
should implement. This allows us to write less coupled code and simplifies
testing.

The third one is to use an interface to declare a type constraint for
generic code.

### Чем отличаются `RWMutex` от `Mutex`?
`Mutex` supports only one kind of `Lock`, whereas `RWMutex` supports
locking for reading (`RLock`/`RUnlock`) and for writing (`Lock`, `Unlock`).

`RWMutex` can have multiple readers but only one writer. When `Lock` is
called (even before the lock is aquired), it blocks `RLock`.

### Чем отличаются буферизированные и не буферизированные каналы?
Buffered channels have non-zero capacity. They can store a number of
values inside. This leads to writing being non-blocking (unless the 
channel buffer is already full). Reading is non-blocking as well (unless
the channel buffer is empty).

Also, the buffer can store values
even after `close` has been called. Of course, writing new values
after closing a channel would lead to a panic. Reading from a closed
channel returns a default value anyway.

### Какой размер у структуры `struct{}{}`?
Size of `struct{}{}` is 0 byte. Every empty struct value is the same object,
so Go doesn't need any more memory for new copies.
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(struct{}{})) // 0
}
```

### Есть ли в Go перегрузка методов или операторов?
No, it was considered to be too much complications for too little value.
Go, however, supports generics, interfaces, variadic functions and struct
embedding that can be used to offer somewhat similar capabilities.

### В какой последовательности будут выведены элементы `map[int]int`?

Maps are unordered in Go, so the order is arbitrary.

### В чем разница `make` и `new`?
`make` allocates memory for and initializes composite data structures, such as
slices, maps and channels. It returns an object itself, not a pointer to it.

`new` returns a pointer to a default value of type. It does not initialize underlying structure of slices, maps and channels (it just returns `nil`
for them).

### Сколько существует способов задать переменную типа `slice` или `map`?

```go
// 5 ways of declaring a slice
var A []int            // nil
B := make([]int, 0)    // initialized, 0 len, 0 cap
C := make([]int, 0, 5) // initialized, 0 len, 5 cap
D := []int{0, 1}

arr := [3]int{0, 1, 2}
E := arr[1:]

// another file
// 4 ways of declaring a map
var A map[int]int         // nil
B := make(map[int]int)    // initialized, small size
C := make(map[int]int, 5) // initialized, enough for 5 ints
D := map[int]int{0: 1, 1: 2}

```

### Что выведет данная программа и почему?

```go
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```
The program prints `1`, `1` because of a mistake in the `update` function.
Developer mixed up *pointer dereference* and *taking a pointer*.

The code can be fixed like this:
```go
func update(p *int) {
	b := 2
	*p = b // fixed line
}
```
Now it prints out `1`, `2`.

### Что выведет данная программа и почему?

```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```
This program is deadlocking because of `wg` copying. `sync.WaitGroup`
contains a `noCopy` field indicating that `sync.WaitGroup` values 
must not be copied after the first use.

The goroutines are using a copy of `wg`, along with a copy of the 
underlying fields, so the invocation of `wg.Done` does not affect the semaphore of "external" `wg` in the main goroutine.

To fix this issue one might either make `wg` argument a pointer, 
or make use of closures by removing the argument altogether.

Then it would print numbers `0` to `4` in arbitrary order and `exit`.

### Что выведет данная программа и почему?

```go
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```
The program prints out `0`, because the variable is shadowed in the
`if` statement, so all the modifications to `n` inside the `if` statement
do not affect external `n`.

This can be fixed by replacing `:=` (declaration operator) with 
`=` (assignment operator). Then the program would print out `2`

### Что выведет данная программа и почему?

```go
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```
The program prints out `[100, 2, 3, 4, 5]` because `v[0] = 100` affects
the underlying array of `v`, which is the same as that of `a`,
but `v = append(v, b)` reassigns `v` to point to a new array, separate from
`a`'s.

To be able to change the array of an external (relative to a function) slice,
the argument must be a pointer and all the internal operations have to be
modified accordingly.

```go
func someAction(v *[]int8, b int8) {
	(*v)[0] = 100
	(*v) = append(*v, b)
}

func main() {
	a := []int8{1, 2, 3, 4, 5}
	someAction(&a, 6)
	fmt.Println(a)
}
```
This program would print out `[100, 2, 3, 4, 5, 6]`.

### Что выведет данная программа и почему?

```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```
The program prints out `[b b a][a a]` because the `slice` variable is
reassigned to a slice different from the external one, so all the
modifications affect the inner `slice` only.

To fix this, the argument has to be made a pointer and all the operations
with `slice` have to be modified accordingly.

```go
func main() {
	slice := []string{"a", "a"}

	func(slice *[]string) {
		*slice = append(*slice, "a")
		(*slice)[0] = "b"
		(*slice)[1] = "b"
		fmt.Print(slice)
	}(&slice)
	fmt.Print(slice)
}
```

Then the program prints out `&[b b a][b b a]`
