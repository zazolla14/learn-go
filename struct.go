package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Merried       bool
}

// ! Struct Method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", customer.Name, ", My Name is", name)
}
func (customer Customer) hu(name string) {
	fmt.Println("Huu", customer.Name, ", My Name is", name)
}

func main() {
	var azola Customer
	azola.Name = "Azola"
	azola.Address = "Cilacap"
	azola.Age = 22

	// fmt.Println(azola)
	// fmt.Println(azola.Name)
	// fmt.Println(azola.Address)
	// fmt.Println(azola.Age)

	// ! Literals
	dety := Customer{
		Name:    "Dety",
		Address: "Cilacap",
		Age:     21,
	}
	// fmt.Println(dety)

	// zola := Customer{"Zola", "Cilacap", 22}
	// fmt.Println(zola)

	// ! Struct Method
	azola.sayHello("Dety")
	dety.hu("Azola")
}
