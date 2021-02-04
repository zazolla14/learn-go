package main

import "fmt"

func endApp() {
	// ! RECOVER
	message := recover()
	if message != nil {
		fmt.Println("(400)Error Message:", message)
	}
	fmt.Println("Ended run Aplication")
}

func runApp(value bool) {
	// !DEFER
	defer endApp()
	fmt.Println("Started  run Application")

	// ! PANIC
	// ! walau error funct endApp akan tetap dijalankan
	// ! panic tidak akan mengeksekusi code dibawahnya dan akan memulai kode ke atas lagi
	if value {
		panic("APPLICATION ERROR")
	}
	fmt.Println("Aplication Running")
}

func main() {
	runApp(true)
	fmt.Println("TEST")
}
