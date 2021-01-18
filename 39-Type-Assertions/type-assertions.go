/**
Type Assertions
1. Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
2. Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
*/

package main

import "fmt"

func random() interface{} {
	return "Budiman"
}

func main() {
	/**
	ketika ingin melakukan type assertions pastikan ketika mengubahnya sesuai dengan
	apa yang diinginkan sesuai function yang sudah dibuat, agar tidak terjadi panic
	dan program akan terhenti dengan sendirinya karena kita salah melakukan type
	assertions
	*/
	result := random()
	resultString := result.(string) // ini merubah type data dari interface kosong menjadi string
	// deklarasinya bisa juga seperti ini
	// var result interface{} = random()
	// var resultString string = result.(string)
	fmt.Println(resultString)

	// ini apabila merubah tipe datanya salah akan menjadi panic, karena return dari function random() ingin dirubah menjadi string
	// resultInt := result.(int) // ini merubah type data dari interface kosong menjadi int
	// fmt.Println(resultInt) // panic, program akan terhenti

	/**
	Kalau ingin aman kita bisa menggunakan

		Type Assertions Menggunakan Switch
		1. Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic
		di aplikasi kita
		2. Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
		3. Agar lebih aman, sebaiknya kita menggunakan switch expression untuk
		melakukan type assertions
	*/
	hasil := random()
	switch value := hasil.(type) { // hasil merupakan value dari func random() yang memiliki interface kosong, jadi bisa tipe data apa saja yang digunakan
	case string: // apabila type datanya string
		fmt.Println("Value", value, "is string")
	case int: // apabila type datanya int
		fmt.Println("Value", value, "is int")
	default: // apabila type data tidak diketahui
		fmt.Println("Unknown Type")
	}
}
