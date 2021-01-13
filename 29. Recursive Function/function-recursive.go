/*
	Recursive Function
	1. Recursive Function adalah function yang memanggil
	function dirinya sendiri
	2. Kadang dalam pekerjaan, kit asering menemui kasus
	dimana menggunakan recursive function lebih mudah
	dibandingkan tidak menggunakan recursive function
	3. Contoh kasus yang lebih mudah diselesaikan menggunakan
	recursive adalah Factorial
*/

package main

import "fmt"

// solusi pertama untuk membuat factorial
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- { // ini menggunakan looping
		// variable i dimulai dari value, value itu parameter function factorialLoop
		result *= i // result = result * i
	}
	return result // kembaliannya adalah nilai dari result
}

// Solusi kedua untuk membuat factorial
func factorialRecursive(value int) int {
	if value == 1 {
		return 1 // kondisi dimana perulangan tersebut berhenti dimana value apabila 1 maka nilai value tersebut adalah 1
	} // kalau tidak 1 maka kita jalankan perintah di bawah
	return value * factorialRecursive(value-1)

}

func main() {
	loop := factorialLoop(10)
	fmt.Println(loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1)

	loopRecursive := factorialRecursive(10)
	fmt.Println(loopRecursive)
}
