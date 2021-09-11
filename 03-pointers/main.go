package main

import "fmt"

func main() {
	var myFavouriteFruit string

	myFavouriteFruit = "apple"

	fmt.Println("my favourite fruit is:", myFavouriteFruit)

	myFavouriteFruit = "pear"

	fmt.Println("my favourite fruit is:", myFavouriteFruit)

	// &myFavouriteFruit -> pointer to a string
	changeFavouriteFruit(&myFavouriteFruit, "peach")

	fmt.Println("my favourite fruit is:", myFavouriteFruit)
	changeFavouriteFruit(&myFavouriteFruit, "banana")
	fmt.Println("my favourite fruit is:", myFavouriteFruit)
}

func changeFavouriteFruit(fruitPointer *string, newFavouriteFruit string) {
	// fmt.Println("the address is:", fruitPointer, "the value is:", *fruitPointer)
	*fruitPointer = newFavouriteFruit
}
