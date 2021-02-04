package main

import "fmt"

func main() {
	name := "Azola"
	counter := 0

	increment := func() {
		name = "budi"
		fmt.Println("INCREMENT")
		counter++
	}

	// * hati2 menggunakan closure
	increment()
	increment()

	fmt.Println(name)
	fmt.Println(counter)
}
