package main

import "fmt"

func main() {
	var number int

	number = 12
	// if statements
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else if number%3 == 0 {
		fmt.Println(number, "is a multiple of 3")
	} else {
		fmt.Println(number, "is odd")
	}

	var favouriteFruit string

	favouriteFruit = "chocolate"

	// switch
	switch favouriteFruit {
	case "apple":
		fmt.Println("yum apples are really nice")
	case "pear":
		fmt.Println("hmmm pears are nice too i think")
	case "banana":
		fmt.Println("you can slip on them, but they're nice")
	default:
		fmt.Println("i don't think that is a fruit")
	}

	// if favouriteFruit == "apple" {
	// 	// do something
	// } else if favouriteFruit == "pear" {
	// 	// do something
	// } else if favouriteFruit == "banana" {
	// 	// do somethign
	// } else {
	// 	// do something
	// }

	// for loops
	i := 1
	for i <= 10 {
		fmt.Println(i, "doubled is:", i+i)
		i = i + 1
	}

	for x := 1; x <= 10; x++ {
		if x%2 == 0 {
			fmt.Println(x, "is even!!")
		}
	}

}
