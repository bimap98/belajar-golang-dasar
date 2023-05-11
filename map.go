package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Eko",
		"city": "Subang",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Eko"
	book["ups"] = "Test"
	fmt.Println(book)
	fmt.Println(book["title"])

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
