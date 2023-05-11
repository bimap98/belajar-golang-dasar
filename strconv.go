package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	decimal := strconv.FormatInt(1000000, 10)
	fmt.Println(decimal)

	value, _ := strconv.Atoi("2000")
	fmt.Println(value)

	valueString := strconv.Itoa(2000)
	fmt.Println(valueString)
}
