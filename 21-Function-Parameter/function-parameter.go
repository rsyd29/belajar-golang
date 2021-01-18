/*
	Function Parameter
	1. Saat membuat function, kadang-kadang kita membutuhkan
	data dari luar, atau kita sebut parameter
	2. Kita bisa menambahkan parameter di function, bisa lebih
	dari satu
	3. Parameter tidaklah wajib, jadi kita bisa membuat function
	tanpa parameter seperti sebelumnya yang sudah kita buat
	4. Namun jika kita menambahkan parameter di function, maka
	kita memanggil function tersebut, kita wajib memasukkan
	data ke parameternya

*/

package main

import "fmt"

/*
- () di dalam kurung buka dan kurung tutup itu merupakan
parameter dari function
- parameter ini bersifat tidaklah wajib, boleh tidak
menggunakannya.
- ketika membuat function itu kadang2 membutuhkan data dari
luar, maka dibuatlah parameter
- untuk menamakan parameter itu mirip seperti menamakan
variabel,
- jadi terdapat nama parameternya, lalu diikuti dengan tipe
datanya
- kalau lebih dari satu parameter, silahkan tambahkan koma
- nama parameter unik, apabila ada 2 parameter atau lebih
maka nama parameter tidak boleh sama, harus berbeda.
- urutannya harus sama ketika ingin memanggil parameternya
pada nama parameter.
*/

func sayHelloTo(firstName string, lastName string) {
	// func yang memiliki 2 parameter
	// nama antar parameter tidak boleh sama,
	// sedangkan tipe data sesuaikan kebutuhannya
	fmt.Println("Hello", firstName, lastName)
}

func main() { // function tanpa parameter, tidaklah wajib
	firstName := "Budiman Rasyid"
	// boleh dimasukan dgn menggunakan variable
	sayHelloTo(firstName, "Zainuddin")
	// urutan nama parameter harus sesuai
	// dan tipe data yang sesuai
	// kalau kita memasukkan kurang atau lebih parameter maka akan terjadi error

	sayHelloTo("Nuratitin", "Suwidi")
	sayHelloTo("Zainuddin", "Fathullah")
	sayHelloTo("Subhan Amin", "Zainuddin")
}
