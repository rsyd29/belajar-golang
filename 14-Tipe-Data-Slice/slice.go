/*
	Tipe Data Slice
	1. Tipe data Slice adalah potongan dari data Array
	2. Slice mirip dengan Array, yang membedakan adalah
	ukuran Slice bisa berubah
	3. Slice dan Array selalu terkoneksi, dimana Slice adalah
	data yang mengakses sebagian atau seluruh data di Array
	(tempat simpan datanya Array, komunikasinya dengan slice)
*/

/*
	Detail Tipe Data Slice
	1. Tipe Data Slice memiliki 3 data, yaitu pointer, length
	dan capacity
	2. Pointer adalah petunjuk data pertama di array para slice
	3. Length adalah panjang dari slice, dan
	4. Capacity adalah kapasitas dari slice, dimana length
	tidak boleh lebih dari capacity
*/

/*
	Membuat Slice dari Array
	Membuat Slice 	--> Keterangan
	array[low:high]	-->	Membuat slice dari array dimulai index low sampai index sebelum high
	array[low:]		--> Membuat slice dari array dimulai index low sampai index akhir di array
	array[:high]	--> Membuat slice dari array dimulai index 0 sampai index sebelum high
	array[:]		--> Membuat slice dari array dimulai index 0 sampai index akhir di array
*/

// Slice itu potongan dari data Array

/*
	Slice itu mengacu pada data Array-nya, apabila data
	Array berubah maka data slice juga akan ikut berubah
*/

package main

import "fmt"

func main() {
	var months = [...]string{ // titik 3 itu berfungsi untuk apabila kita tidak mengetahui berapa kapasitasnya
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	/*
		Function Slice
		Operasi		--> Keterangan
		len(slice)	--> Untuk mendapatkan panjang
		cap(slice)	--> Untuk mendapat kapasitas
		append(slice, data)	--> Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
		make([]TypeData, length, capacity) --> Membuat slice baru
		copy(destination, source) --> Menyalin slice dari source ke destination
	*/

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// kalau kita ubah data arraynya maka slice juga akan ikut berubah
	// months[5] = "Data Array Diubah"
	// fmt.Println(slice1)

	// kalau kita ubah data slice maka array juga akan ikut berubah
	// slice1[0] = "Data Slice Index 0 Diubah"
	// fmt.Println(months)

	// Kode Program Append Slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Ahad"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Ahad Baru"
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jum'at Sabtu Baru Ahad Baru]

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Budiman") // Berfungsi untuk membuat Array baru, apabila sanggup menampungnya, jika tidak maka akan membuat Array baru, ketika di buat array baru lalu ubah data slice nya yang kena dampak adalah Array barunya, Array lama tidak akan berubah
	// tidak merubah existing Array aslinya, karena dia akan membuat Array baru
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2) // [November Desember]
	fmt.Println(months) // [Januari Februari Maret April Mei Juni Juli Agustus September Oktober November Desember]

	/*
		Make Slice
	*/
	newSlice := make([]string, 2, 5) //paling aman ketika ingin membuat slice dari awal, karena Array nya ada di slice jadi tidak terlihat
	newSlice[0] = "Budiman"
	newSlice[1] = "Rasyid"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	/*
		Copy Slice, pastikan ukurannya sama kalau tidak nanti akan terpotong
	*/
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	/*
		NOTE: HATI-HATI Saat Membuat Array
		1. Saat membuat Array, kita harus berhati-hati, jika
		salah, maka yang kita buat bukanlah Array, melainkan
		Slice

		Array vs Slice
	*/

	iniArray := [...]int{1, 2, 3, 4, 5} // harus masukkan panjang dari Array tersebut di dalam []
	iniSlice := []int{1, 2, 3, 4, 5}    // jangan diberikan angka di []
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
