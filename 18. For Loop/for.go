/*
	Perulangan For Loops (cuman 1 di golang)
	1. Dalam bahasa pemrograman, biasanya ada fitur yang
	bernama perulangan
	2. Salah satu fitur perulangan adalah for loops
*/

package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		// kalau kondisi true maka akan menjalankan statement yang ada di kurung kurawal
		fmt.Println("Perulangan ke ", counter)
		counter++ // ini sama seperti counter = counter + 1
		// ketika false yaitu 11 maka akan berhenti programnya
	}

	/*
		For dengan statement
		1. Dalam for, kita bisa menambahkan statement, dimana
		terdapat 2 statement (init dan post) yang bisa
		ditambahkan di for
		2. Init statement, yaitu statement sebelum for di
		eksekusi
		3. Post statement, yaitu statement yang akan selalu
		dieksekusi di akhir tiap perulangan
	*/

	// kode program di atas bisa disingkat menjadi seperti ini
	for counterBaru := 1; counterBaru <= 10; counterBaru++ {
		// for initStatement; kondisi; postStatement
		fmt.Println("Perulangan ke", counterBaru)
	}

	/*
		For Range
		1. For bisa digunakan untuk melakukan iterasi terhadap
		semua data collection
		2. Data collection contohnya Array, Slice dan Map
	*/

	// Slice ingin di initialisasikan semuanya
	slice := []string{"Budiman", "Rasyid", "Zainuddin", "Nuratitin", "Subhan", "Amin"}
	// cara manualnya seperti ini
	fmt.Println("Cara Manual")
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
	fmt.Println(slice[4])
	fmt.Println(slice[5])

	// cara lamanya seperti ini
	fmt.Println("Cara lama Iterasi Slice")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Cara cepat kalau mau mengambil data yang banyak di Array/Slice/Map
	fmt.Println("\nCara cepat Iterasi Slice/Array/Map dengan for range")
	for _, value := range slice {
		// fmt.Println("Index", i, "=", value)
		fmt.Println(value) // apabila variable i tidak terpakai maka kita memakai _ (underscore) untuk menggantikan i-nya
	}

	fmt.Println("\nCara cepat Iterasi Slice/Array/Map dengan for range pakai index")
	// kalau slice/array berupa index
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
		// fmt.Println(value) // apabila variable i tidak terpakai maka kita memakai _ (underscore) untuk menggantikan i-nya
	}

	person := make(map[string]string)
	person["name"] = "Budiman Rasyid"
	person["title"] = "Programmer"
	// kalau map berupa key
	for key, value := range person {
		fmt.Println("\n", key, "=", value)
	}
}
