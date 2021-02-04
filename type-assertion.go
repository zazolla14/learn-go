package main

import "fmt"

func Ok() interface{} {
	return "Azola"
}

func main() {
	result := Ok()
	switch dataType := result.(type) {
	case string:
		fmt.Println("String", dataType, "is 'string'")
	case int:
		fmt.Println("integer", dataType, "is 'int'")
	default:
		fmt.Println("unknow type")
	}

	// result := Ok()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)
}
