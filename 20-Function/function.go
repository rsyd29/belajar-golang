/*
	Function
	1. Sebelumnya kita sudah mengenal sebuah function yang
	wajib dibuat agar program kita bisa berjalan, yaitu function
	main
	2. Function adalah sebuah blok kode program yang sengaja
	dibuat dalam program agar bisa digunakan berulang-ulang
	3. Cara membuat function sangat sederhana, hanya dengan
	menggunakan kata kunci func lalu diikuti dengan nama function
	nya dan blok kode isi function nya
	4. Setelah membuat function, kita bisa mengeksekusi function
	tersebut dengan memanggilnya menggunakan kata kunci nama
	function nya diikuti tana kurung buka, kurung tutup
*/

package main

import "fmt"

func sayHello() { // best practice untuk membuat function yaitu camle case sayHello
	fmt.Println("Hello World!")
}

func main() {
	for i := 0; i < 10; i++ {
		sayHello() // memanggil function hanya menuliskan nama function tersebut
		// di golang bersifat case sensitive artinya huruf besar dan kecil berpengaruh
		/*
			ini merupakan func sayHello yang sedang dipanggil
			di func main, maka func main ini akan memanggil func sayHello
			sebanyak looping for tersebut
		*/
	}
}
