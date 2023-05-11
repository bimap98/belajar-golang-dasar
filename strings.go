package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("  Eko Kurniawan Khannedy   ", " "))
	fmt.Println(strings.ToLower("Eko Kurniawan Khannedy"))
	fmt.Println(strings.ToUpper("Eko Kurniawan Khannedy"))
	fmt.Println(strings.Split("Eko Kurniawan Khannedy", " "))
	fmt.Println(strings.Contains("Eko Kurniawan Khannedy", "Eko"))
	fmt.Println(strings.Contains("Eko Kurniawan Khannedy", "Budi"))
	fmt.Println(strings.ReplaceAll("Eko Kurniawan Khannedy", "Eko", "Budi"))

}
