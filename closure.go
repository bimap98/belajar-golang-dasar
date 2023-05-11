package main

import "fmt"

func main() {

	name := "Eko"
	counter := 0

	increment := func() {
		name = "Tono"
		counter++
	}

	increment()

	fmt.Println(name)
	fmt.Println(counter)

}
