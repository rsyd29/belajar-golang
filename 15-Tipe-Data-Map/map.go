/*
	Tipe Data Map
	1. Pada Array atau Slice, Untuk mengakses data, kita
	menggunakan index Number dimulai dari 0
	2. Map adalah tipe data lain yang berisikan kumpulan
	data yang sama, namun kita bisa menentukan jenis tipe
	data index yang akan kita gunakan
	3. Sederhananya, Map adalah tipe data kumpulan key-value
	(kata kunci-nilai), dimana kata kuncinya bersifat unik,
	tidak boleh sama
	4. Berbeda dengan Array dan Slice, jumlah data yang kita
	masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan
	kata kuncinya berbeda, jika kita gunakan kata kunci sama,
	maka secara otomatis data sebelumnya akan diganti dengan
	data baru
*/

package main

import "fmt"

func main() {
	person := map[string]string{
		//"key": "value"
		"name":    "Budiman",
		"address": "JakSel",
	}

	// Untuk mengubah/menambahkan data di Map
	person["title"] = "Programmer"

	fmt.Println(person) // untuk menampilkan seluruh Map
	// untuk mengaksesnya seperti ini
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	/*
		Function Map
		Operasi --> Keterangan
		len(map) --> Untuk mendapatkan jumlah data di map
		map[key] --> Mengambil data di map dengan key
		map[key] = value --> Mengubah data di map dengan key
		make(map[TypeKey]TypeValue) --> Membuat map baru
		delete(map, key) --> Menghapus data di map dengan key
	*/

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Budiman"
	book["ups"] = "Error Nih"
	fmt.Println(book)
	fmt.Println(len(book)) // panjang sesuai dengan isi dari data mapnya, berbeda dengan array dan slice dia sesuai panjang arraynya

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
