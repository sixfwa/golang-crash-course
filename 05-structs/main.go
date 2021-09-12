package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Country   string
}

type Book struct {
	Title  string
	Author string
}

func main() {
	personOne := Person{
		FirstName: "Frodo",
		LastName:  "Baggins",
		Country:   "Shire",
	}

	bookOne := Book{
		Title:  "Lord of the Rings",
		Author: "Tolkien",
	}

	fmt.Println("personOne:", personOne.getFullName())

	fmt.Println("bookOne title:", bookOne.Title)
	bookOne.changeTitle("Harry Potter")
	fmt.Println("bookOne title:", bookOne.Title)

}

// receiver / method
func (book *Book) changeTitle(newTitle string) {
	book.Title = newTitle
}

// receiver / method
func (person *Person) getFullName() string {
	return person.FirstName + " " + person.LastName
}
