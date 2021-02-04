package main

import "fmt"

func Ups(value int) interface{} {
	if value == 1 {
		return 1
	} else if value == 2 {
		return false
	} else {
		return "Ups"
	}
}

func main() {
	result := Ups(3)
	fmt.Println(result)
}
