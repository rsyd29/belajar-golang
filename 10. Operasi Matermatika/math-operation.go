/*
	Operasi Matematika
	Operator --> Keterangan
	+ --> Penjumlahan
	- --> Pengurangan
	* --> Perkalian
	/ --> Pembagian
	% --> Sisa Pembagian (Modulus)
*/

package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	/*
		Augmented Assigments
		(Untuk melakukan operasi matematika terhadap var
		dia sendiri)
		Operasi Matematika --> Augmented Assignments
		a = a + 10 --> a += 10
		a = a - 10 --> a -= 10
		a = a * 10 --> a *= 10
		a = a / 10 --> a /= 10
		a = a % 10 --> a %= 10
	*/

	var i = 100
	i += 100 // ini adalah i = i + 10

	fmt.Println(i)

	/*
		Unary Operator
		(Operasi ke satu buah var)
		Operator --> Keterangan
		++ --> a = a + 1
		-- --> a = a - 1
		- --> Negative
		+ --> Positive
		! --> Boolean kebalikan
	*/

	i++ // artinya i = i + 1
	fmt.Println(i)

	var negative = -100
	var positive = +100 // tanpa + juga bisa by default
	fmt.Println(negative)
	fmt.Println(positive)
}
