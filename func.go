package main

import "fmt"

// ! basic
// func sayHello(name string) string {
// 	if name == "" {
// 		return "Hello Wolrd!"
// 	} else {
// 		return "Hello " + name + "!"
// 	}
// }

// ! multyple values
// func getFullName() (string, string, int8) {
// 	return "Azola", "Zubizarreta", 22
// }

// ! named return values
// func getFullName() (firstName, lastName string, age int8) {
// 	firstName = "Azola"
// 	lastName = "Zubizarreta"
// 	age = 22

// 	return
// }

// ! Variadic Func
func sumall(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func main() {
	// ! basic
	// result := sayHello("Azola")
	// fmt.Println(result)
	// fmt.Println(sayHello(""))

	// ! multyple values
	// * gunakan _ untuk meng ignore values yang tidak ingin digunakan
	// firstName, _, age := getFullName()
	// fmt.Println(firstName)
	// //  fmt.Println(lastName)
	// fmt.Println(age)

	// ! named return values
	// * gunakan _ untuk meng ignore values yang tidak ingin digunakan
	// _, lastName, age := getFullName()
	// // fmt.Println(firstName)
	// fmt.Println(lastName)
	// fmt.Println(age)

	// ! Variadic Func
	result := sumall(10, 10, 10, 10, 10)
	fmt.Println(result)

	slice := []int{10, 20, 30, 40, 50}
	total := sumall(slice...)
	fmt.Println(total)
}
