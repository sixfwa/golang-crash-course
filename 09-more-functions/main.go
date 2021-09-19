package main

import "fmt"

func main() {
	fmt.Println("the average:", calculateAverage(5.5, 4.5, 6.7, 9.8))
	someStuff := []string{"asdkjfh", "asfha", "sadf"}
	fmt.Println(someStuff, 234.3, 654, "asdf")
	boringFunction("nice world", 324.9, "covid sucks", 3234)

	func() {
		fmt.Println("my internal function")
	}()

	add := func(numberOne, numberTwo int) int {
		return numberOne + numberTwo
	}

	fmt.Println("using the internal add():", add(34, 535))
	fmt.Println("using the internal add():", add(3, 5))

	hundred := 100
	fmt.Println("big number:", hundred)
	addHundred := func() int {
		hundred += 100
		return hundred
	}

	fmt.Println("big number:", addHundred())
	fmt.Println("big number:", addHundred())
}

func boringFunction(someBoringStuff ...interface{}) {
	fmt.Println(someBoringStuff...)
}

func calculateAverage(numbers ...float64) float64 {
	fmt.Println(numbers)
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}
