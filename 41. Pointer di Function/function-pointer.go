/**
Pointer di Function pada Golang
1. Saat kita membuat parameter di function, secara default adalah pass by value,
artinya data akan di copy lalu dikirim ke function tersebut
2. Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya
tidak akan pernah berubah
3. Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
4. Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter
tersebut
5. Untuk melakukan ini, kita juga bisa menggunakan pointer di function
6. Untuk menjadikan sebuah parameter menjadi pointer, kita bisa menggunakan operator
* di parameternya
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address Address) {
	address.Country = "Indonesia" // data duplikat
}

func ChangeCountryToIndonesiaPointer(address *Address) {
	address.Country = "Indonesia" // data duplikat
}

func main() {
	var alamat = Address{ // data asli
		City:     "Cirebon",
		Province: "Jawa Barat",
		Country:  "",
	}
	fmt.Println(alamat)

	ChangeCountryToIndonesia(alamat) // alamat ini bukan pointer, ini data aslinya
	fmt.Println(alamat)              // tidak berubah data aslinya, karena data itu di copy ke dalam parameter func ChangeCountryToIndonesia()

	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesiaPointer(alamatPointer) // alamat menjadi pointer tambahkan operator &
	// code di bawah sama aja seperti code di atas
	// ChangeCountryToIndonesiaPointer(&alamat) // alamat menjadi pointer tambahkan operator &
	fmt.Println(alamat) // ini berubah dengan data pointer yang di duplikat

	/**
	Saran, kalau memiliki data struct yang banyak yaitu fieldnya, jadikan pointer saja
	kenapa? tiap memanggil function maka  data tersebut akan selalu di duplikat di memory
	jadi, semakin banyak manupalis pemanggilan function yang memanfaatkan parameter struct itu
	maka memory itu akan semakin bengkak overload, lebih baik menggunakan pointer
	dibanding menggunakan value
	*/
}
