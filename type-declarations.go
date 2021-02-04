package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var AzolaKTP NoKTP = "123123213123"
	var AzolaMerried Married = false
	fmt.Println(AzolaKTP)
	fmt.Println(NoKTP("321321321321"))
	fmt.Println(AzolaMerried)
}
