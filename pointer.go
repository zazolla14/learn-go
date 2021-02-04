package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// ! Pointer Function
func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Jepang"
}

func main() {
	//! Passing by Value, mengcopy value nya ke variabel baru
	address1 := Address{
		City:     "Cilacap",
		Province: "Jawa Tengah",
		Country:  "Indeonesia",
	}

	// address2 := address1
	// address2.City = "Bogor"
	// fmt.Println(address1) // * address1 tidak ikut berubah karena passing by value, bukan by references
	// fmt.Println(address2)

	// ! Pass by Reference dengan Pointer &, memaksa address2 mengukiti variabel address2
	// address2 := &address1
	// address2.City = "Bogor"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// !Pointer *, untuk membuat pointer, mengubah variabel apapun yang mengacu ke Address struck
	address2 := &address1 //* artinya address1 akan ikut berubah jika address2 diubah
	address3 := Address{
		City:     "London Tengah",
		Province: "London",
		Country:  "England",
	}
	address4 := address3

	// address2 = &Address{
	// 	City:     "Maos",
	// 	Province: "Madrid",
	// 	Country:  "Spanyol",
	// }

	// fmt.Println(address1) // ! tidak akan mengubah address1 karena address2 akan membuat data baru
	// fmt.Println(address2)

	address4.City = "Kroya"

	*address2 = Address{ // ? semua variable yang emnggunakan struct Address akan berubah datanya
		City:     "Maos",
		Province: "Madrid",
		Country:  "Spanyol",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)

	// ! New, membuat pointer dengan menggunakan Pointer &
	// ! mengembalkikan pointer kosong
	var alamat1 *Address = new(Address)
	alamat2 := alamat1

	alamat2.City = "Mernek"
	fmt.Println(alamat1)
	fmt.Println(alamat2)

	// ! Pointer Function
	address := Address{
		City:     "Kesugihan",
		Province: "Jawa Tengah",
		Country:  "",
	}
	ChangeAddressToIndonesia(&address) // ! membuat pointer &, artinya address akan mengikuti func
	fmt.Println(address)               // ! tidak akan berubah, karena data dari address yang dikirim adalah data copy
}
