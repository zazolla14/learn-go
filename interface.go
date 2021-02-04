package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name

}
func main() {
	var azola Person
	azola.Name = "Azola"
	fmt.Println(azola.Name)
	sayHello(azola)

	name := azola.GetName()
	fmt.Println(name)

	person := Person{Name: "Azola"}
	sayHello(person)

	macan := Animal{
		Name: "macan",
	}

	sayHello(macan)
}
