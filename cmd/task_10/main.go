package main

import "fmt"

/*
Дана последовательность температурных колебаний:
	-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

// temperature group step
const BucketSize = 10

func main() {
	temperature := []float32{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5,
	}

	groups := make(map[int][]float32)
	for _, t := range temperature {
		group := int(t) / BucketSize * BucketSize
		if values, ok := groups[group]; !ok {
			// if the group does not exist yet
			groups[group] = []float32{t}
		} else {
			// if the group is already there
			groups[group] = append(values, t)
		}
	}

	fmt.Println(groups)
}
