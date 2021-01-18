/*
	Closures
	1. Closure adalah kemampuan sebuah function berinteraksi
	dengan data-data disekitarnya dalam scope yang sama
	2. Harap gunakan fitur closure ini dengan bijak saat kita
	membuat aplikasi
*/

package main

import "fmt"

func main() {
	// scope --> lingkup kerja sebuah variable/function
	// biasanya menggunakan blok kurung kurawal {}
	// artinya apabila ada sebuah function/variable berada pada kurung kurawal itu maka akan bisa berinteraksi dengan data yang ada pada kurung kurawal tersebut
	name := "Budiman"
	counter := 0
	increment := func() { // ini anonymous function
		// kalau kita membuat function di dalam function maka data tersebut hanya bisa berinteraksi dengan yang ada blok kurung kurawal
		fmt.Println("increment")
		counter++ // karena var counter itu berapa pada scope function main maka kita bisa memasukkan ke dalam anonymous function tersebut
		// var atas bisa di akses dengan yang ada pada didalamnya
		// sedangkan var bawah tidak bisa di akses pada bagian atas
		// counter++ ini merupakan sebuah closure dimana kita bisa merubah data yang ada pada scope tersebut yaitu variable counter yang nilai awalnya 0, bisa dirubah karena counter++
		// kalau kita menggunakan perintah di bawah ini
		// name = "rasyid"  // maka var name yang ada di atas ikut berubah
		name := "rasyid" // apabila menggunakan ini dia membuat variable baru di anonymous function ini, jadi variabel name yang di atas tidak akan berubah
		fmt.Println(name)
		// jadi fitur closure itu dapat merubah isi data yang berada di atas, akan tetapi harus hati hati
	}

	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
