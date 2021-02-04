package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	nilai64 := int64(nilai32)
	nilai16 := int16(nilai32) // * jika nilai int sudah melewati maksimal maka akan kembali ke nilai minimal, int16 maksimal nilai dari -32767 sampai 32768 jika menampung lebih dari 32768 makan value akan dihitung kembali dari -32767
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Azola"
	var e byte = name[0] // * e bertipe data uint8 atau byte
	eString := string(e)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
