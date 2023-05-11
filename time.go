package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.YearDay())

	utc := time.Date(2010, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	parse, _ := time.Parse("2006-01-02", "2020-10-15")
	fmt.Println(parse)

}
