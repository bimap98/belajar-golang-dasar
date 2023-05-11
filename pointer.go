package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	*address2 = Address{"Subang", "Jawa Barat", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	address4 := new(Address)
	address4.City = "Surabaya"

	fmt.Println(address4)

	alamat := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)

	fmt.Println(alamat)

}
