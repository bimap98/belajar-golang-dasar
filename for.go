package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Counter ke ", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Ini counter ke ", counter)
	}

	slice := []string{"Eko", "Kurniawan", "Khannedy"}

	//for i := 0; i < len(slice); i++ {
	//	fmt.Println(slice[i])
	//}

	for _, value := range slice {
		//fmt.Println("Index ke ", i, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)

	person["name"] = "Eko"
	person["city"] = "Subang"

	for key, value := range person {
		fmt.Println("Key =", key, "Value =", value)
	}

}
