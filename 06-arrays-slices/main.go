package main

import (
	"fmt"
	"sort"
)

func main() {
	// Array of 4 elements
	var sweets = [4]string{
		"milk chocolate", // index 0
		"dark chocolate", // index 1
		"marshmellow",    // index 2
		"toffee apple",   // index 3
	}

	fmt.Println("sweet at index 0:", sweets[0])
	sweets[1] = "caramel"
	fmt.Println("sweet at index 1:", sweets[1])
	fmt.Println("length of sweets:", len(sweets))

	// Slice
	var scores []int

	scores = append(scores, 10)
	scores = append(scores, 14)
	scores = append(scores, 12)
	scores = append(scores, 18)

	fmt.Println("scores slice:", scores)

	sort.Ints(scores)

	fmt.Println("scores slice:", scores)

	// for i := 0; i < len(scores); i++ {
	// 	fmt.Println("index:", i, "value:", scores[i])
	// }

	for _, score := range scores {
		fmt.Println("score:", score)
	}

	fmt.Println("the average score is:", averageScore(scores))
	// scores[4] = 18 cannot do
}

func averageScore(scores []int) float64 {
	total := 0
	for _, score := range scores {
		total += score
	}

	return float64(total / len(scores))
}
