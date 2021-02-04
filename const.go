package main

import "fmt"

func main() {
	const firstName string = "Azola"
	const lastName = "Zubizarreta"
	fmt.Println(firstName)
	fmt.Println(lastName)

	// ! ERROR
	// firstName = "Dety"
	// lastName = "Mei"

	const (
		country = "Indonesia"
		distric = "Cilacap"
	)
	fmt.Println(country)
	fmt.Println(distric)
}
