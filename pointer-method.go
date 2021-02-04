package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	azola := Man{
		Name: "Azola",
	}
	azola.Married()
	fmt.Println(azola.Name)
}
