/*
	Interface Kosong
	1. Go-Lang bukanlah bahasa pemrograman yang berorientasi
	objek
	2. Biasanya dalam pemrograman berorientasi objek, ada satu
	data parent di puncak yang bisa dianggap sebagai semua
	implementasi data yang ada di bahasa pemrograman tersebut
	3. Contoh di Java ada java.lang.Object
	4. Untuk menangani kasus seperti ini, di Go-Lang kita bisa
	menggunakan interface kosong
	5. Interface kosong adalah interface yang tidak memiliki
	deklarasi method satupun, hal ini membuat secara otomatis
	semua tipe data akan menjadi implementasi nya
*/

/*
	Penggunaan Interface Kosong di Go-Lang
	* Ada banyak contoh penggunaan interface kosong di Go-Lang, seperti:
	isinya interface kosong, oleh karena itu bisa memasukkan apapun di dalam function tersebut
	! fmt.Println(a ...interface{})
	! panic(v interface{})
	! recover() interface{}
	! dan lain-lain
*/

package main

import "fmt"

func Ups(i int) interface{} {
	/*
		NOTE: jadi kita bisa return tipe data apapun, karena bergantung pada kontrak interface kosong tersebut
		ini tidak akan error karena interface itu bisa menjadi tipe data apapun
		jadi seakan akan kontrak di golang tipe data apapun masuk ke dalam interface kosong
	*/
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	kosong := Ups(1) // ini tipe datanya merupakan interface kosong
	var data interface{} = Ups(3)
	fmt.Println(kosong)
	fmt.Println(data)
}
