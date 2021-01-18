/*
	Type Declarations
	1. Type Declarations adalah kemampuan membuat ulang tipe
	data baru dari tipe data yang sudah ada
	2. Type Declarations biasanya digunakan untuk membuat
	alias terhadap tipe data yang sudah ada, dengan tujuan
	agar lebih mudah dimengerti
*/

package main

import "fmt"

func main() {
	type NoKTP string
	// ini untuk membuat alias dengan tipe data string
	type Married bool
	// ini untuk membuat alias dengan tipe data boolean

	var ktpBudiman NoKTP = "1111111111111111"
	var marriedStatus Married = true
	fmt.Println(ktpBudiman)
	fmt.Println(NoKTP("222222222222222"))
	fmt.Println(marriedStatus)
}
