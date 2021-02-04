package main

import "fmt"

func main() {
	value := 0

	for value <= 10-1 {
		fmt.Println("Loops", value+1)
		value++
	}

	for value := 1; value <= 10; value++ {
		fmt.Println("Perulangan", value)
	}

	slice := []string{"Azola", "Zubi", "Zarreta"}
	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Azola"
	person["country"] = "Indonesia"
	fmt.Println(person)

	for key, v := range person {
		fmt.Println("Key", key, "=", v)
	}
}
