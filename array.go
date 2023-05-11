package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"

	fmt.Println(names[0])

	var value = []int{
		10,
		20,
		30,
	}
	value[0] = 11
	fmt.Println(value[0])
	fmt.Println(len(value))

	var lagi [10]string
	fmt.Println(len(lagi))
}
