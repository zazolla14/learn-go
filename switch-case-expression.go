package main

import "fmt"

func main() {
	name := "Azola Zubizarreta"
	fmt.Println(len(name))

	switch name {
	case "Joko":
		fmt.Println("Error")
	case "Azola":
		fmt.Println("Oke")
	default:
		fmt.Println("Error")
	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("error")
	case false:
		fmt.Println("Oke")
	}

	length := len(name)
	switch {
	case length <= 5:
		fmt.Println("error")
	case length <= 25:
		fmt.Println("Oke")
	}
}
