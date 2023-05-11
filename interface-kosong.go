package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return "Satu"
	} else if i == 2 {
		return 2
	} else {
		return true
	}
}

func main() {

	data := Ups(3)
	fmt.Println(data)

}
