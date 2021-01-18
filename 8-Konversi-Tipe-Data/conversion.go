/*
	Konversi Tipe Data
	1. Di Go-Lang kadang kita butuh melakukan konversi tipe
	data dari satu tipe ke tipe lain
	2. Misal kita ingin mengkonversi tipe data int32 ke int64,
	dan lain-lain atau sebaliknya
*/

package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32) // ini akan ke int64
	var nilai16 int16 = int16(nilai64) // ini akan ke int16
	var nilai8 int8 = int8(nilai32)    // terjadi int overflow
	/*
		^^ hati hati kalau ukurannya tidak cukup, terjadi
		perubahan value
	*/

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	/*
		Kode program konversi tipe data (2)
	*/

	var name = "Budiman Rasyid"
	var B byte = name[0] // ini byte itu alias dari uint8
	var eString string = string(B)
	/*
		kode diatas digunakan untuk konversi dari byte ke
		string
	*/
	fmt.Println(name)
	fmt.Println(B)
	fmt.Println(eString)
}
