/*
	Struct
	1. Struct adalah sebuah template data yang digunakan untuk
	menggabungkan nol atau lebih tipe data lainnya dalam satu
	kesatuan
	(jadi bisa lebih dari satu tipe data)
	2. Struct biasanya representasi data dalam program aplikasi
	yang kita buat
	3. Data di struct disimpan dalam field
	4. Sederhananya struct adalah kumpulan dari field
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

	// untuk merepresentasikan sebuah data diusahakan menggunakan struct
	// karena lebih baik strukturnya dibandingkan dengan Map ataupun Array

	fmt.Println(budiman)
	fmt.Println(budiman.Name)
	fmt.Println(budiman.Address)
	fmt.Println(budiman.Age)

	/*
		Struct Literals
		1. Sebelumnya kita telah membuat data dari struct,
		namun sebenarnya ada banyak cara yang bisa kita
		gunakan untuk membuat data dari struct

		berikut kode program struct literals
	*/

	// lebih baik deklarasikan seperti ini
	rasyid := Customer{
		Name:    "Rasyid",
		Address: "Setiabudi",
		Age:     21,
	}
	fmt.Println(rasyid)

	// ingat deklarasi struct nya harus sama penempatannya
	// akan tetapi kalau pakai ini tidak disarankan
	// karena apabila type struct kita berubah maka
	// kode dibawah akan error
	zainuddin := Customer{"Zainuddin", "Anggrek", 53}
	fmt.Println(zainuddin)
}
