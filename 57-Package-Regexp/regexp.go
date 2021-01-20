/**
Package Regexp (Regular Expression)
1. Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
2. Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
3. Ini untuk mempelajari syntaxnya https://github.com/google/re2/wiki/Syntax/
4. Dokumentasinya adalah https://golang.org/pkg/regexp/

Beberapa Function di Package regexp
Function --> Kegunaan
regexp.MustCompile(string) --> Membuat Regexp, string disitu membuat pola regexp
regexp.MatchString(string) bool --> Mengecek apakah Regexp match dengan string
regexp.FindAllString(string, max) --> Mencari string yang match dengan maximum jumlah hasil
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("([a-z])([a-z])([a-z])([a-z])man")

	fmt.Println(regex.MatchString("budiman"))
	fmt.Println(regex.MatchString("Budiman"))
	fmt.Println(regex.MatchString("rohman"))
	fmt.Println(regex.MatchString("utsman"))
	fmt.Println(regex.MatchString("burhman"))

	// ini akan mencari string yang ada pada variable regex yang sudah kita initial
	search := regex.FindAllString("budiman bumaman buja budeman buka bula bola rohman", -1) // ini stringnya akan di cari sebanyak n return value slice of string
	// -1 itu tidak ada batasnya untuk mencarinya
	fmt.Println(search)

	// regexp berguna untuk melakukan validasi dan apapun yang berhubungan dengan teks
	// pelajari di dokumentasinya
}
