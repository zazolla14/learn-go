package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) { // blacklist(name) akan menjalan func yang dikirim dari mainb() dan mencocokan name dengan "admin"
		fmt.Println("You are block", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklistAdmin := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklistAdmin)
	registerUser("Azola", func(name string) bool {
		return name == "root"
	})
}
