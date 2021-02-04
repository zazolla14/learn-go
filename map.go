package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Azola",
		"address": "Cilacap",
	}

	person["profession"] = "BE Engeneer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Basic Go"
	book["author"] = "Azola"
	book["wrong"] = "wrong"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")
	fmt.Println(book)
	fmt.Println(len(book))
}
