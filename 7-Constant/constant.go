/*
Constant
1. Constant adalah variable yang nilainya tidak bisa diubah lagi setelah
pertama kali diberi nilai
2. Cara pembuatan cpnstant mirip dengan variable, yang membedakan hanya kata
kunci yang digunakan adalah const, bukan var
3. Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya
*/

package main

import "fmt"

func main() {
	const firstName string = "Budiman"
	const lastName = "Rasyid"
	const value = 1000

	//fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// akan error
	//firstName = "Tidak Bisa Diubah"
	//lastName = "Tidak Bisa Diubah"

	/*
	 yang membedakan antara const dan var adalah kalau const apabila tidak
	 digunakan maka tidak ada error

	 kebutuhan const itu untuk dipakai diberbagai source code

	 kalau var dan tidak digunakan maka akan terjadi error
	*/

	/*
		kalau const tidak bisa menggunakan := kalau pakai itu maka akan
		menjadi var, kalau const pakai =
	*/

	/*
		Deklarasi Multiple Constant
		1. Sama seperti var, di Go-Lang juga kita bisa membuat constant secara
		sekaligus banyak
	*/
	const (
		country string = "Indonesia"
		city           = "Setiabudi"
		number         = 27
	)

	fmt.Println(country)
	fmt.Println(city)
	fmt.Println(number)
}
