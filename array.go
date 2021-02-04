package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Azola"
	names[1] = "Zubi"
	names[2] = "Zarreta"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	names[2] = "Dety"
	fmt.Println(names[2])
	fmt.Println(names)

	var values = [3]int{
		10, 20, 30,
	}
	fmt.Println(len(values))
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var test [10]string
	fmt.Println("maksimal aray =", len(test)) // * yang dibaca maksumal isi array bukan jumlah isi array
}
