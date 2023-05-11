package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}

func main() {
	eko := NewMap("Eko")
	fmt.Println(eko)
}
