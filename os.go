package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument: ")

	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	value, err := os.Hostname()
	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println("Error", err.Error())
	}

	javaHome := os.Getenv("JAVA_HOME")
	fmt.Println(javaHome)
}
