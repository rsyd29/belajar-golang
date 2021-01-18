/*
	Operasi Perbandingan
	1. Operasi perbandingan adalah operasi untuk membandingkan
	dua buah data (semua data di golang bisa dibandingkan)
	2. Operasi perbandingan adalah operasi yang menghasilkan
	nilai boolean (benar atau salah)
	3. Jika hasil operasinya adalah benar, maka nilainya
	adalah true
	4. Jika hasil operasinya adalah salah, maka nilainya
	adalah false

	Operator Perbandingan
	Operator --> Keterangan
	> --> Lebih Dari
	< --> Kurang Dari
	>= --> Lebih Dari Sama Dengan
	<= --> Kurang Dari Sama Dengan
	== --> Sama Dengan
	!= --> Tidak Sama Dengan
*/

package main

import "fmt"

func main() {
	var name1 = "Budiman"
	var name2 = "Budi"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	/*
		Berikut adalah operasi perbandingan di golang
	*/
	fmt.Println(value1 > value2)  // lebih besar dari
	fmt.Println(value1 < value2)  // lebih kecil dari
	fmt.Println(value1 == value2) // sama dengan
	fmt.Println(value1 != value2) // tidak sama dengan
}
