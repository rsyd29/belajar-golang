/**
Package Strings
1. Package strings adalah package yang berisikan function-function untuk memanupulasi
tipe data string
2. Ada banyak sekali function yang bisa kita gunakan
3. Dokumentasinya adalah https://golang.org/pkg/strings/

Beberapa Function di Package strings
Function --> Kegunaan
strings.Trim(string, cutset) --> Memotong cutset di awal dan akhir string (menghilangkan spasi yang ada pada sisi kiri dan kanan)
strings.ToLower(string) --> Membuat semua karakter string menjadi lower case
strings.ToUpper(string) --> Membuat semua karakter string menjadi upper case
strings.Split(string, separator) --> Memotong string berdasarkan separator (bertujuan untuk split ke slice dengan spasi)
strings.Contains(string, search) --> Mengecek apakah string mengandung string lain
strings.ReplaceAll(string, from, to) --> Mengubah semua string dari from ke to
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Budiman Rasyid Zainuddin", "Zainuddin")) // return value bool
	// output true, karena string Zainuddin ada pada string
	fmt.Println(strings.Contains("Budiman Rasyid Zainuddin", "Nuratitin")) // return value bool
	// output false, karena string Nuratitin tidak ada pada string

	fmt.Println(strings.Split("Budiman Rasyid Zainuddin", " ")) // return value []string slice string
	// output [Budiman Rasyid Zainuddin], karena di split menjadi slice of string

	fmt.Println(strings.ToLower("Budiman Rasyid Zainuddin")) // return value string
	// output budiman rasyid zainuddin, karena function ToLower mengubah string menjadi huruf kecil semua
	fmt.Println(strings.ToUpper("Budiman Rasyid Zainuddin")) // return value string
	// output BUDIMAN RASYID ZAINUDDIN, karena function ToUpper mengubah string menjadi huruf besar semua
	fmt.Println(strings.ToTitle("budiman rasyid zainuddin")) // return value string
	// output hasilnya sama seperti ToUpper

	fmt.Println(strings.Trim("     Budiman Rasyid Zainuddin     ", " ")) // return value string
	// output Budiman Rasyid Zainuddin, karena ada spasi di kiri dan kanan (biasanya berguna untuk ketika user melakukan input data)
	fmt.Println(strings.Trim("a     Budiman Rasyid Zainuddin     a", " ")) // return value string
	// output a     Budiman Rasyid Zainuddin     a, karena ada huruf a di kiri dan kanannya jadi spasinya tetap tidak akan hilang
	fmt.Println(strings.ReplaceAll("Budiman Budiman Budiman", "Budi", "Uts")) // return value string
	// output Utsman Utsman Utsman, karena kata Budi diganti menjadi Uts, oleh karena itu outputnya Utsman Utsman Utsman
}
