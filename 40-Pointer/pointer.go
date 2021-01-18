/**
Pass by Value
1. Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
2. Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
sebenarnya yang dikirim adalah duplikasi value nya
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//address1 := Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"} // ini merupakan struct
	//address2 := address1 // ini di assign, jadi data address 1 itu di duplikat ke address2
	//
	//address2.City = "Jakarta Pusat" // lalu kita ubah address2 Citynya menjadi Cirebon
	//fmt.Println(address1) // address1 tidak berubah, karena data address 1 di duplikat ke address2
	//fmt.Println(address2)

	/**
	Penjelasan Detail Pass by Value
	jadi data dari address1 itu di copied ke address2
	(jadi di memory itu ada 2 lokasi yang berbeda untuk ke dua data tersebut)
	*/

	/**
	Pointer
	1. Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa
	menduplikasi data yang sudah ada
	2. Sederhananya, dengan kemampuan pointer, kita bisa membuat Pass by Reference
	*/

	/**
	Pass by Reference dengan Pointer
	Ketika kita membuat address2 dari address1, maka dia akan mengacu pada memory yang sama
	yaitu memory address1, apabila kita mengubah data dari address2 maka data dari addres1
	juga ikut berubah
	*/

	/**
	Untuk melakukan Pointer, kita menggunakan Operator &
	1. Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa
	menggunakan operator & diikuti dengan nama variable nya
	*/

	// kode dibawah sama seperti ini
	// var alamat1 Address = Address{"Cirebon", "Jawa Barat", "Indonesia"}
	address1 := Address{"Cirebon", "Jawa Barat", "Indonesia"}

	/**
	Operator *
	1. Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
	2. Semua variable yang mengacu ke data yang sama tidak akan berubah
	3. Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa
	menggunakan operator *
	*/

	// kode dibawah sama seperti ini
	// var alamat2 *Address = &alamat1
	address2 := &address1 // penggunaan pointer seperti ini
	// berguna untuk alamat2 mengambil data dari alamat1
	// jadi alamat2 itu akan reference ke alamat1

	address2.City = "Sukabumi" // ini data alamat2 untuk City diubah menjadi Sukabumi
	// maka alamat1 dengan City yang awalnya Cirebon akan berubah menjadi Sukabumi
	// kenapa? karena alamat2 merupakan pointer ke alamat1

	// bagaimana kalau kita assign alamat2 ke pointer address yang lainnya
	// alamat2 = &Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"}
	// kode diatas tidak merubah alamat1 nya, jadi sekarang alamat2 yang jadi pointer
	// jadi ini membuat memory baru dengan
	// nama alamat2 yang isinya struct Address di atas

	// bagaimana caranya agar yang awalnya alamat1 itu mengacu pada memory alamat1 lalu dirubah ke memory alamat2?
	// solusinya dengan Operator *
	*address2 = Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"}

	// kita coba buat alamat3
	var address3 *Address = &address1 // apakah ikut berubah atau tidak? dia ikut berubah
	//var address4 *Address = &Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"} // ini pointer

	/**
		Function New
	1. Sebelumnya untuk membuat pointer dengan menggunakan operator &
	2. Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
	3. Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
	*/

	var address4 *Address = new(Address) // ini akan menjadi data kosong
	// lalu kita coba ubah
	address4.City = "Cirebon"
	// bisa juga seperti ini untuk membuat baru
	// var address4 *Address = &Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"} // ini pointer

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
}
