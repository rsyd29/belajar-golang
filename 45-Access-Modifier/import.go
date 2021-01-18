package main

import (
	"belajar-golang/45-Access-Modifier/helper"
	"fmt"
)
import "belajar-golang/45-Access-Modifier/other"

// nama foldernya harus sesuaikan dengan gopathnya
// ~/go/src/belajar-golang

func main() {
	helper.SayHello("Budiman")
	// helper.sayGoodBye("Budiman") // ini error karena func sayGoodBye diawali dengan huruf kecil
	// oleh karena itu tidak bisa di akses dari luar package
	other.SayHello("Rasyid")

	fmt.Println(helper.Application) // ini bisa di akses karena di awali dengan huruf besar
	// fmt.Println(helper.version) // ini tidak bisa di akses karena di awali dengan huruf kecil
}
