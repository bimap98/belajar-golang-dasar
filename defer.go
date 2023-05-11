package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil aplikasi")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Hasil", result)
}

func main() {
	runApplication(0)
}
