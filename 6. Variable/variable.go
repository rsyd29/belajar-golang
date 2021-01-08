/*
1. Variable adalah tempat untuk menyimpan data
2. Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita
mau
3. Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita
ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa
variable
4. Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti
dengan nama variable dan tipe datanya
*/

package main

import "fmt"

func main() {
	var name string
	// di golang kalo variablenya tidak dipakai maka itu dianggap error

	name = "Budiman Rasyid"
	fmt.Println(name)

	name = "Budiman Rasyid Zainuddin"
	fmt.Println(name)

	/*
		Tipe Data Variable
		1. Saat kita membuat variable, maka kita wajib menyebutkan tipe data
		variable tersebut
		2. Namun jika kita langsung meninisialisasikan data pada variablenya,
		maka kita tidak wajib menyebutkan tipe data variablenya
	*/

	var fullName = "Budiman Rasyid"
	fmt.Println(fullName)

	fullName = "Budiman Rasyid Zainuddin"
	fmt.Println(fullName)

	var age = 20 // ini tipe datanya int32 atau int64 tergantung OS-nya
	fmt.Println(age)

	/*
		Kata Kunci Var
		1. Di Go-Lang, kata kunci var saat membuat variable tidaklah wajib
		2. Asalkan saat membuat variable kita langsung menginisialisasi datanya
		3. Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan
		kata kunci := saat menginisialisasikan data pada variable tersebut
	*/

	country := "Indonesia"
	fmt.Println(country)

	/*
		Deklarasi Multiple Variable
		1. Di Go-Lang kita bisa membuat variable secara sekaligus banyak
		2. Code yang dibuat akan lebih bagus dan mudah dibaca
	*/

	var (
		firstName = "Budiman Rasyid"
		lastName  = "Zainuddin"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
