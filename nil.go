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
	person := NewMap("Azola")
	fmt.Println(person)

	data := NewMap("")

	if data == nil {
		fmt.Println("Empety Data")
	} else {
		fmt.Println(data)
	}

	// var person map[string]string = nil
	// fmt.Println(person)
}
