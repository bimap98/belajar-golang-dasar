package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eTo"))

	fmt.Println(regex.FindAllString("eko eto edo ePo", 2))
	fmt.Println(regex.FindAllString("eko eto edo ePo", -1))
}
