package main

import "fmt"

func main() {

	var name = "Eko"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Kurniawan" {
		fmt.Println("Hello Kurniawan")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Panjang karakter lebih dari lima")
	} else {
		fmt.Println("Panjang karakter kurang dari lima")
	}

}
