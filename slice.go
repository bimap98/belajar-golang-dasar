package main

import "fmt"

func main() {

	months := [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	slice := months[2:4]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//months[2] = "Maret Lagi"
	//months[3] = "April Lagi"

	slice2 := months[9:]
	slice3 := append(slice2, "Eko")
	fmt.Println(slice3)
	slice3[0] = "November Baru"
	fmt.Println(slice3)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
