/*
	Anonymous Function
	1. Sebelumnya setiap membuat function, kita akan selalu
	memberikan sebuah nama pada function tersebut
	2. Namun kadang ada kalanya lebih mudah membuat function
	secara langsung di variable atau parameter tanpa harus
	membuat function terlebih dahulu
	3. Hal tersebut dinamakan anonymous function, atau function
	tanpa nama
*/

package main

import "fmt"

/*
ini merupaka type declaration dengan nama Blacklist yang
merupakan sebuah function dengan parameter string dan return
value berupa bool
*/
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	/*
		ini merupakan function registerUser
		yang memiliki 2 parameter yaitu
		name dengan tipe data string
		blacklist dengan sebuah type declaration Blacklist
	*/

	if blacklist(name) {
		// apabila nilainya true maka perintah bawah dijalankan
		fmt.Println("You Are Blocked", name)
	} else {
		// apabila false
		fmt.Println("Welcome", name)
	}
}

// sebelumnya kita buat function ini dulu

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

// kita biasanya males membuat function lebih baik kita langsung deklarasikan dengan variablenya saja
// caranya seperti ini
func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	/* kode diatas merupakan anonymous function karena tidak
	ada nama functionnya, melainkan hanya sebuah nilai dari
	variable blacklist
	*/

	registerUser("admin", blacklist)
	registerUser("rasyid", blacklist)
	// kita juga bisa langsung menambahkan ke dalam parameter
	registerUser("budiman", func(name string) bool {
		return name == "root"
	})

	registerUser("root", func(name string) bool {
		return name == "root"
	})
}
