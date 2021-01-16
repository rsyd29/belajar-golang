/**
Error Interface
1. Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error,
nama interface nya adalah error

the error built-in interface type is the conventional interface for representing
an error condition, with the nil value representing no error

type error interface { // di golang sendiri sudah ada
	Error() string // membuat function Error() dengan return value nya adalah string
}
*/

package main

import (
	"fmt"
	"github.com/pkg/errors"
)

/**
 	Membuat Error
	1. Untuk membuat error, kita tidak perlu manual
	2. Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang
	terdapat di package errors (Package akan dibahas secara detail di materi sendiri)
*/
func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError)

	// ini memiliki 2 variable yang memiliki value berupa func Pembagi dengan 2 parameter
	hasil, err := Pembagi(100, 0)
	if err == nil { // jika var err == nil, artinya tidak error maka
		fmt.Println("Hasil", hasil)
	} else { // jika punya error maka kita akan print errornya apa
		fmt.Println("Error", err.Error())
	}
}
