package helper

import (
	"fmt"
)

var version = "1.0.0"              // tidak bisa di akses dari luar package, karena di awali dengan huruf kecil
var Application = "Belajar Golang" // bisa di akses dari luar package, karena di awali dengan huruf besar

func SayHello(name string) { // // bisa di akses dari luar, karena di awali dengan huruf besar
	fmt.Println("Hello", name)
	sayGoodBye("Budiman")
	fmt.Println(version)
}

func sayGoodBye(name string) { // karena di awali dengan huruf kecil maka tidak bisa digunakan oleh package lain
	// artinya tidak bisa diakses dari luar package
	fmt.Println("Goodbye", name)
}
