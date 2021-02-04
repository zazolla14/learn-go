package main

import (
	"fmt"
	"learn-go/script/helper"
	"learn-go/script/other"
)

func main() {
	result := helper.SayHello("azola")
	fmt.Println(result)

	result2 := other.SayHello("Dety")
	fmt.Println(result2)
}
