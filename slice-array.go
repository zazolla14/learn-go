package main

import "fmt"

func main() {
	months := [...]string{
		"Januari", "February", "March", "April", "May", "Juny", "July", "Agust", "September", "October", "November", "Desember",
	}
	slice1 := months[4:7]
	slice2 := months[6:9]
	fmt.Println(months)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println((cap(slice1)))
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice1[0] = "dety"
	fmt.Println(slice1)
	fmt.Println(slice1[0])

	days := [...]string{
		"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu",
	}
	slice3 := days[5:]
	fmt.Println(days)
	fmt.Println(slice3)

	slice3[0] = "Sabtu Baru"
	slice3[1] = "Minggu Baru"
	fmt.Println(slice3)

	appendSlice4 := append(slice3, "Hari Hari")
	fmt.Println(appendSlice4)
	appendSlice4[0] = "Sabtu Sabtu"
	fmt.Println(appendSlice4)
	fmt.Println(days)

	makeSlice := make([]string, 2, 5)
	makeSlice[0] = "Azola"
	makeSlice[1] = "Dety"
	fmt.Println(makeSlice)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))

	rescSlice := makeSlice[:]
	fmt.Println(rescSlice)
	destSlice := make([]string, len(rescSlice), cap(rescSlice))
	copy(destSlice, rescSlice)
	fmt.Println(destSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
