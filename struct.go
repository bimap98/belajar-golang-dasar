package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi", customer.Name, "my name", name)
}

func (a Customer) sayHuu() {
	fmt.Println("Huu", a.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Subang"
	eko.Age = 30

	eko.sayHi("Joko")
	eko.sayHuu()

	//fmt.Println(eko)
	//
	//budi := Customer{
	//	Name:    "Budi",
	//	Address: "Cirebon",
	//	Age:     35,
	//}
	//
	//fmt.Println(budi)
	//
	//joko := Customer{"Joko", "Jakarta", 30}
	//
	//fmt.Println(joko)
}
