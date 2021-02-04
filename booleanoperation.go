package main

import "fmt"

func main() {
	ujian := 80
	absensi := 80
	// var lulusUjian = ujian >= 80
	// var lulusabsensi = absensi > 80

	// var lulus = lulusUjian && lulusabsensi
	// fmt.Println(lulus)

	fmt.Println("lulus =", ujian >= 80 && absensi > 80)
	fmt.Println("lulus =", !(ujian >= 80 && absensi > 80))
}
