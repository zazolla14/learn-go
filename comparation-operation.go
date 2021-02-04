package main

import "fmt"

func main() {
	name1 := "Azola"
	name2 := "Dety"
	var result bool = name1 == name2
	fmt.Println(result)

	value1 := 100
	value2 := 100
	fmt.Println(value1 > value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
