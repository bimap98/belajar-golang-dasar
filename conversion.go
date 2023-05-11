package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai16 = int16(nilai32)
	var nilai8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	var name = "Eko"
	var eName = name[0]
	var eString = string(eName)
	fmt.Println(name)
	fmt.Println(eString)
}
