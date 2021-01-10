/*
	If Expression
	1. If adalah salah satu kata kunci yang digunakan untuk
	percabangan
	2. Percabangan artinya kita bisa mengeksekusi kode program
	tertentu ketika suatu kondisi terpenuhi
	(Jika kondisi terpenuhi maka akan dieksekusi, dan jika
	tidak terpenuhi maka tidak akan dieksekusi)
	3. Hampir di semua bahasa pemrograman mendukung if expression
*/

package main

import "fmt"

func main() {
	name := "Budi"

	if name == "Budiman" { // ini harus boolean apabila true maka akan dieksekusi dalam kurung kurawal

		// Ini di dalam kurung kurawal akan dieksekusi apabila true
		fmt.Println("Hello Budiman")
		/*
			Else Expression
			1. Blok if akan dieksekusi ketika kondisi if bernilai
			true
			2. Kadang kita ingin melakukan eksekusi program tertentu
			jika kondisi if bernilai false
			3. Hal ini bisa dilakukan menggunakan else expression
		*/

		/*
			Else If Expression
			1. Kadang dalam If, kita butuh membuat beberapa
			kondisi
			2. Kasus seperti ini, kita bisa menggunakan
			Else If Expression
		*/
	} else if name == "Rasyid" {
		// ini kondisi percabangan kedua apa bila true maka akan mengeksekusi kode program dibawah ini
		// apabila kondisi banyak kita tinggal menambahkan else if yang lainnya
		fmt.Println("Hello Rasyid")
	} else {
		// ini apabila kondisi bernilai false akan dieksekusi
		fmt.Println("Hi, Boleh Kenalan Kaga?")
	}

	/*
		If dengan Short Statement (ini hanya ada di golang)
		1. If mendukung short statement sebelum kondisi
		2. Hal ini sangat cocok untuk membuat statement yang
		sederhana sebelum melakukan pengecekan terhadap kondisi
	*/

	//biasanya seperti ini
	lengthLama := len(name)
	if lengthLama > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	// kalau di golang bisa seperti ini, lebih simple
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
