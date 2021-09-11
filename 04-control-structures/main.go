package main

import "fmt"

func main() {
	var number int

	number = 10

	// if statements
	if number%2 == 0 {
		if number == 10 {
			fmt.Println("woah it's:", number)
		}
		fmt.Println(number, "is even")
	} else if number%3 == 0 {
		fmt.Println(number, "is a multiple of 3")
	} else {
		fmt.Println(number, "is odd")
	}

	var myFruit string

	myFruit = "chocolate"

	// switch statements
	switch myFruit {
	case "apple":
		fmt.Println("that's also my favourite fruit")
	case "pear":
		fmt.Println("pears are okey in the morning")
	case "banana":
		fmt.Println("you can slip on it")
	case "peach":
		fmt.Println("they are nice in the summer")
	default:
		fmt.Println("im not sure that's a fruti")
	}

	// for loop
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println("i is:", i)
		}
		i = i + 1
	}

	for x := 1; x <= 20; x++ {
		if x%3 == 0 {
			fmt.Println(x, "is a number")
		}
	}

}
