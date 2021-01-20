/**
Package Reflect (Reflection)
1. Dalam bahasa pemrograman, biasanya ada fitur reflection, dimana kita bisa melihat
struktur kode kita pada saat aplikasi sedang berjalan
2. Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
3. Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, Anda bisa
eksplorasi package reflect ini secara otodidak
4. Reflection sangat berguna ketika kita ingin membuat library yang general sehingga
mudah digunakan
5. Dokumentasinya adalah https://golang.org/pkg/reflect/
*/

package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	/**
	Kode Program Struct Tag
	berguna untuk menambahkan tag ke dalam informasi struct, cocok untuk membuat library
	untuk melakukan validasi, dibandingkan cek manual menggunakan if else, lebih baik
	menggunakan automation package reflection ini

	Dengan menggunakan ini untuk melakukan validasi, dan buatlah framework validasi
	*/
	Name string `required:"true" max:"10"`
	// required -> nama tag
	// true -> value tag
	// ketika ingin menambahkan lebih dari satu tag, tambahkan spasi
	// jadi struct tag di atas terdapat 2 tag
	Age int
}

type ContohLagi struct { // ini tidak mungkin ke validasi karena tidak ada tag required nya
	Name        string `required:"true"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data) // mengambil data reflection ke dalam var t
	// melakukan iterasi semua field nya
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// lalu di cek apakah fieldnya memiliki tag dengan nama required dan valuenya adalah true
		if field.Tag.Get("required") == "true" {
			// kalau true, akan di cek datanya, apabila datanya itu sama dengan string kosong, maka itu tidak valid
			value := reflect.ValueOf(data).Field(i).Interface() // artinya ini tidak boleh sampai kosong
			if value == "" {
				return false // karena return false, jadi tidak valid
			}
			// atau bisa juga seperti ini kodenya
			// return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	// kalau semuanya tidak kosong baru dianggap valid
	return true
}

func main() {
	sample := Sample{"Budiman", 20}
	sampleType := reflect.TypeOf(sample)        // return value object type
	structField := sampleType.Field(0)          // ini field ke index 0
	required := structField.Tag.Get("required") // untuk mengambil tag required

	fmt.Println(sampleType.NumField()) // berguna untuk ada berapa jumlah field-nya
	fmt.Println(structField.Name)      // ini untuk ingin tahu nama fieldnya apa
	fmt.Println(required)
	fmt.Println(sampleType.Field(0).Tag.Get("max")) // untuk mengambil tag max
	fmt.Println(sampleType.Field(0).Tag.Get("min")) // apabila tagnya tidak ada, outputnya string kosong

	fmt.Println(IsValid(sample)) // apabila datanya ada, maka valid
	sample.Name = ""
	fmt.Println(IsValid(sample)) // apabila datanya string kosong, maka tidak valid

	contoh := ContohLagi{"Budiman", "Oke"}
	fmt.Println(IsValid(contoh)) // maka ini hasilnya true, walaupun itu string kosong, solusinya tambahkan struct tag pada struct ContohLagi

}
