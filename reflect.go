package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `require:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("require") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}

	}
	return true
}

func main() {
	sample := Sample{"Eko"}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)

	fmt.Println(sampleType.Field(0).Tag.Get("require"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	fmt.Println(IsValid(sample))
}
