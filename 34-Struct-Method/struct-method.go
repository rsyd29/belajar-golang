/*
	Struct Method
	1. Struct adalah tipe data seperti tipe data lainnnya,
	dia bisa digunakan sebagai parameter untuk function
	2. Namun jika kita ingin menambahkan method ke dalam
	structs, sehingga seakan-akan sebuah struct memiliki
	function
	3. Method adalah function
*/

package main

import "fmt"

// penamaan struct biasanya setiap awal kata huruf besar
type Customer struct {
	// layaknya memberikan field pada table customer
	// penamaan field juga sama setiap awal kata huruf besar
	Name, Address string
	Age           int
}

// ini merupakan struct method/function
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

// ini merupakan struct method/function
func (a Customer) sayHuuu() {
	fmt.Println("Huuuuu from", a.Name)
}

/*
	Membuat Data Struct
	1. Struct adalah template data atau prototype data
	2. Struct tidak bisa langsung digunakan
	3. Namun kita bisa membuat data/object dari struct yang
	telah kita buat
*/

func main() {
	// cara penggunaan struct seperti ini
	var budiman Customer
	budiman.Name = "Budiman"
	budiman.Address = "Indonesia"
	budiman.Age = 21

	// lalu panggilanya seperti ini
	budiman.sayHi("Rasyid")
	budiman.sayHuuu()

}
