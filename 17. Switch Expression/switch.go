/*
	Switch Expression
	1. Selain if expression, untuk melakukan percabangan,
	kita juga bisa mengggunakan Switch Expression
	2. Switch Expression sangat sederhana dibandingkan if
	3. Biasanya switch expression digunakan untuk melakukan
	pengecekkan ke kondisi dalam satu variable
	(versi sederhananya dari if expression)
*/

package main

import "fmt"

func main() {
	name := "Budi" // variable awal

	switch name { // variable awal yang akan kita beri suatu kondisi
	case "Budiman": // ini sebuah kondisinya
		fmt.Println("Hello Budiman") // ini statement apabila kondisi terpenuhi
	case "Rasyid":
		fmt.Println("Hello Rasyid")
	default: // kondisi apabila tidak ada yang terpenuhi (ini seperti else)
		fmt.Println("Kamu siapa ya?") // maka akan menjalankan statement ini
	}

	/*
		Switch dengan short statement
		1. Sama dengan if, switch juga mendukung short statement
		sebelum variable yang akan di cek kondisinya
	*/
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
		// kenapa tidak ada default karena hasilnya hanya dua kondisi yaitu true or false (boolean)
	}

	/*
		Switch Tanpa Kondisi
		1. Kondisi di switch expression tidak wajib
		2. Jika kita tidak menggunakan kondisi di switch expression,
		kita bisa menambahkan kondisi tersebut di setiap case nya
	*/
	lengthName := len(name)
	switch {
	// lebih mirip dengan if tapi lebih sederhana, akan tetapi kalau kondisinya lebih kompleks maka tetap pakai if expression
	case lengthName > 10: // ini kondisinya di dalam case
		fmt.Println("Nama Terlalu Panjang")
	case lengthName > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
