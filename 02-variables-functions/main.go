package main

import "fmt"

var GlobalVariable = "this is a global variabel"

func main() {
	var awesomeString string // declaration of a variable
	var reallyCoolInteger int

	awesomeString = "hello awesome world" // assignment of awesomeString
	reallyCoolInteger = 23
	coolFloat := 23.4 // declaration + assigment

	fmt.Println("awesomeString is:", awesomeString)
	fmt.Println("reallyCoolInteger is:", reallyCoolInteger)
	fmt.Println("coolFloat is:", coolFloat)
	fmt.Println("globalVariable", GlobalVariable)

	// firstName := getFirstName()
	fmt.Println("firstName is", getFirstName())

	_, carol := getABunchOfNames()

	fmt.Println("name two:", carol)

	incrementedNumber := incrementByOne(11)
	fmt.Println("incrementedNumber:", incrementedNumber)
	ourNewFloat := addTwoFloats(13.3, 45.3)
	fmt.Println("ourNewFloat:", ourNewFloat)

}

func SubtractTwoIntegers(numberOne, numberTwo int) int {
	return numberOne - numberTwo
}

func addTwoFloats(numberOne, numberTwo float64) float64 {
	return numberOne + numberTwo
}

func incrementByOne(number int) int {
	return number + 1
}

func getFirstName() string {
	return "rithmic"
}

func getABunchOfNames() (string, string) {
	return "Bob", "Carol"
}
