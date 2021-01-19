/**
package strconv (string conversion)
1. Sebelumnya kita sudah belajar cara konversi tipe data, misal dari in32 ke int64
number ke number lain
2. Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? misal dari int
ke string, atau sebaliknya dan selainnya
3. Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
4. Documentasinya adalah https://golang.org/pkg/strconv/

Beberapa Function di Package strconv
Function --> Kegunaan
strconv.parseBool(string) --> Mengubah string ke bool
strconv.parseFloat(string) --> Mengubah string ke float
strconv.parseInt(string) --> Mengubah string ke Int64
strconv.FormatBool(bool) --> Mengubah bool ke string
strconv.FormatFloat(float, ...) --> Mengubah float64 ke string
strconv.FormatInt(int, ...) --> Mengubah int64 ke string

jadi parse itu tipe data string ke tipe data yang akan dituju
sedangkan Format itu tipe data yang lain menjadi string
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// di function strconv itu dia return valuenya ada 2
	// 1. datanya dan 2. errornya
	boolean, err := strconv.ParseBool("true")
	if err == nil { // apabila errornya itu tidak ada maka
		fmt.Println(boolean)
	} else { // apabila terdapat error maka
		fmt.Println(err)
	}

	// ParseInt(string, base, bitSize)
	// string yang ingin dirubah
	// base itu ke mana desimal (10), biner (2), oktal(8), heksadesimal(16)
	// bitSize itu ukuran dari tipe datanya
	number, err := strconv.ParseInt("1000000", 10, 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err)
	}

	// return valuenya hanya satu, datanya langsung jadi
	biner := strconv.FormatInt(1000000, 2)
	oktal := strconv.FormatInt(1000000, 8)
	desimal := strconv.FormatInt(1000000, 10)
	hexa := strconv.FormatInt(1000000, 16)
	fmt.Println(biner)
	fmt.Println(oktal)
	fmt.Println(desimal)
	fmt.Println(hexa)

	// cara paling gampang dari pada ParseInt
	valueInt, err := strconv.Atoi("2000000") // string ke int
	if err == nil {
		fmt.Println(valueInt)
	} else {
		fmt.Println(err)
	}

	valueString := strconv.Itoa(1000000) // int ke string
	fmt.Println(valueString)

}
