package main

import "fmt"

type User struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	grades := make(map[string]int)

	grades["Harry Potter"] = 23
	grades["John Doe"] = 34
	grades["Frodo Baggins"] = 30

	fmt.Println("grades:", grades)
	delete(grades, "Frodo Baggins")
	fmt.Println("grades:", grades)

	userDatabase := make(map[string]*User)

	userOne := User{FirstName: "Michael", LastName: "Jackson", Age: 54}
	userTwo := User{FirstName: "Albus", LastName: "Dumbledore", Age: 70}
	userThree := User{FirstName: "Jon", LastName: "Snow", Age: 34}

	userDatabase["1"] = &userOne
	userDatabase["2"] = &userTwo
	userDatabase["3"] = &userThree

	fmt.Println("user database:", userDatabase)
	userOne.FirstName = "Jermaine"
	fmt.Println(userDatabase["1"].FirstName, userOne.FirstName)

	for key, value := range userDatabase {
		fmt.Println("user id:", key, "user:", value.FirstName)
	}
}
