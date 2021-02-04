package main

import "fmt"

func main() {
	var name string
	name = "Azola"
	fmt.Println(name)
	name = "Zubizarreta"
	fmt.Println(name)

	var age = 30
	fmt.Println(age)
	age = 22
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)
	country = "America"
	fmt.Println(country)

	var (
		firstName = "Azola"
		lastName  = "Zubizarreta"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	firstName = "Dety"
	lastName = "Dina"
	fmt.Println(firstName)
	fmt.Println(lastName)
}
