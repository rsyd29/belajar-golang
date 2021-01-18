/*
1. String ada tipe data kumpulan karakter
2. Jumlah karakter di dalam String bisa nol sampai tidak terhingga
3. Tipe data String di Go-Lang direpresentasikan dengan kata kunci string
4. Nilai data String di Go-Lang selalu diawali dengan karakter " (petik dua)
dan diakhiri dengan karakter " (petik dua) lagi
*/

package main

import "fmt"

func main() {
	fmt.Println("Budiman")
	fmt.Println("Budiman Rasyid")
	fmt.Println("Budiman Rasyid Zainuddin")

	/*
		Function untuk String

		1. len("string") --> Menghitung jumlah karakter di String
		2. "string"[number] --> Mengambil karakter pada posisi yang ditentukan
		number itu index (dimulai dari 0)
	*/

	fmt.Println(len("Budiman"))
	fmt.Println("Budiman Rasyid"[0])           //outputnya 66 itu kode ASCII dari B
	fmt.Println("Budiman Rasyid Zainuddin"[1]) //ouputnya 117 itu kode ASCII dari u

}
