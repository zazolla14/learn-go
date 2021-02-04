package main

import "fmt"

func main() {
	name := "Azola"
	if name == "Joko" {
		fmt.Println("Salah lah")
	} else if name == "Bowo" {
		fmt.Println("Salah lagi lah")
	} else {
		fmt.Println("Betul")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nam terlalu panajng")
	} else {
		fmt.Println("Oke")
	}
}
