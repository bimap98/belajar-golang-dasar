package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)
	//resultInteger := result.(int)
	//fmt.Println(resultInteger)

	switch value := result.(type) {
	case string:
		fmt.Println("Value nya string", value)
	case int:
		fmt.Println("Value nya int", value)
	default:
		fmt.Println("Unknown")

	}

}
