package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put yout host")
	user := flag.String("user", "root", "Put yout user")
	password := flag.String("password", "root", "Put yout password")
	number := flag.Int("number", 100, "Put yout password")

	flag.Parse()

	fmt.Println("Host", *host)
	fmt.Println("User", *user)
	fmt.Println("Password", *password)
	fmt.Println("Number", *number)

}
