package main

import "fmt"

func getFullNameTwo() (firstName, middleName, lastName string) {
	return "Eko", "Kurniawan", "Khannedy"
}

func main() {
	a, b, c := getFullNameTwo()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
