package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian > 75
	var lulusAbsensi = absensi > 85

	var hasil = lulusUjian || lulusAbsensi
	fmt.Println(hasil)

	fmt.Println(!true)
}
