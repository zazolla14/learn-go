package main

import (
	"errors"
	"fmt"
)

func Pembagian(value int, value2 int) (int, error) {
	if value2 == 0 {
		return 0, errors.New("Value2 cannot use 0")
	} else {
		return value / value2, nil
	}
}

func main() {
	result, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Result :", result)
	} else {
		fmt.Println("Result", result)
		fmt.Println("Error", err.Error())
	}
}
