/*
	Function as Parameter
	1. Function tidak hanya bisa kita simpan di dalam
	variable sebagai value
	2. Namun juga bisa kita gunakan sebagai parameter untuk
	function lain

	Jadi function yang kita gunakan sebagai value di sebuah
	variable kita dapat parsing ke dalam function lain
	sebagai parameter

	Ini merupakan fitur yang sangat powerfull
*/

package main

import "fmt"

/*
	Function Type Declarations
	1. Kadang jika function terlalu panjang, agak ribet untuk
	menuliskannya di dalam parameter
	2. Type Declarations juga bisa digunakan untuk membuat
	alias function, sehingga akan mempermudah kita menggunakan
	function sebagai parameter
*/

type Filter func(string) string // ini merupakan type declarations agar nama parameternya tidak terlalu panjang

func sayHelloWithFilter(name string, filter Filter) { // jadi pada paramater ke 2 kita tidak harus menyebutkan func(string) string cukup nama alias type declarationnya saja
	// diatas merupakan function dengan 2 parameter
	// ada 1 parameter yang bertipe function yaitu filter yang mengembalikan data dengan tipe data string
	nameFiltered := filter(name) // ini merupakan sebuah variable yang nilainya itu sebuah function filter dengan parameter name yang ada pada function parameter sayHelloWithFilter
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		// kata kasar Anjing kita ubah nilainya menjadi ...
		return "..."
	}
	return name // kalau bukan maka nilainya name itu sendiri
}

func main() {
	sayHelloWithFilter("Budiman", spamFilter) // ini parameter langsung menggunakan nama functionnya
	filter := spamFilter                      // function spamFilter sebagai value dari variable filter
	sayHelloWithFilter("Anjing", filter)      // parameter ke 2 itu merupakan sebuah variable yang valuenya itu function spamFilter

}
