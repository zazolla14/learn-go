package main

import "fmt"

func main() {
	var a = 10
	b := 10
	var c = a * b
	fmt.Println(c)

	result := 10 + 10
	fmt.Println(result)

	var i = 10
	i %= 2
	fmt.Println(i)

	i++ // = 0+1 = 1
	i++ // = 1+1 = 2
	fmt.Println(i)

	positive := +14
	negative := -14
	fmt.Println(positive)
	fmt.Println(negative)

}
