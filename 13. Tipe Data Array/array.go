/*
	Tipe Data Array
	1. Array adalah tipe data yang berisikan kumpulan data
	dengan tipe yang sama
	2. Saat Membuat Array, kita perlu menentukan jumlah data
	yang bisa ditampung oleh Array tersebut
	3. Daya tampung Array tidak bisa bertambah setelah Array
	dibuat
*/

/*
Index di array di awali angka 0, yang berisikan tipe data yang sama
saat membuat array harus menentukan jumlah data yang bisa ditampung
Daya tampung tidak bisa bertambah setelah array dibuat,'
kalau mau menambahkan data harus membuat array baru.
*/

package main

import "fmt"

func main() {
	//Array secara manual
	var names [3]string

	names[0] = "Budiman"
	names[1] = "Rasyid"
	names[2] = "Zainuddin"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//Array langsung isi datanya tanpa ribet
	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	/*
		Function di Array
	*/

	fmt.Println(len(names)) //Panjang Arraynya bukan jumlah arraynya
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi)) // output 10

}
