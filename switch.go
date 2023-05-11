package main

import "fmt"

func main() {

	name := "Eo"

	switch name {
	case "Eko":
		fmt.Println("Halo Eko")
	case "Joko":
		fmt.Println("Halo Joko")
	default:
		fmt.Println("Kenalan Donk")
	}

	//switch length := len(name); length >= 3 {
	//case true:
	//	fmt.Println("Panjang karakter lebih atau sama dengan tiga")
	//case false:
	//	fmt.Println("Panjang karakter kurang dari tiga")
	//}

	length := len(name)

	switch {
	case length >= 3:
		fmt.Println("Panjang karakter lebih atau sama dengan tiga")
	case length < 3:
		fmt.Println("Panjang karakter kurang dari tiga")
	default:
		fmt.Println("Ini Default")
	}

}
