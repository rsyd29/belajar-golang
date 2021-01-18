/*
	Returning Multiple Values (Di Go-Lang)
	1. Function tidak hanya dapat mengembalikan satu value,
	tapi juga bisa multiple value
	2. Untuk memberitahu jika function mengembalikan multiple
	value, kita harus menulis semua tipe data return value nya
	di function
*/

package main

import "fmt"

/*
	Menghiraukan Return Value
	1. Multiple return value wajib ditangkap semua value
	nya
	2. Jika kita ingin menghiraukan return value tersebut,
	kita bisa menggunakan tanda _ (garis bawah)
*/

func getFullName() (string, string, string) {
	// return value-nya itu kalau lebih dari 2 itu harus pakai tanda (tipe data return valuenya)
	return "Budiman", "Rasyid", "Zainuddin"
	// Budiman itu return value 1
	// Rasyid itu return value 2
	// Zainuddin itu return value 3
}

func main() {
	firstName, middleName, _ := getFullName() // ini bisa simpan di 2 variable
	// firstName itu mengambil nilai dari return value 1
	// middleName itu mengambil nilai dari return value 2
	// lastName itu mengambil nilai dari return value 3
	fmt.Println(firstName, middleName)
	fmt.Println(firstName)
	fmt.Println(middleName)
	// fmt.Println(lastName)
	// Jadi kita bisa simpan return lebih dari satu data

}
