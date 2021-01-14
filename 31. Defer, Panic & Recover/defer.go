/*
	Defer
	1. Defer Function adalah function yang bisa kita jadwalkan
	untuk dieksekusi setelah sebuah function selesai di eksekusi
	2. Defer Function akan selalui dieksekusi walaupun terjadi
	error di function yang di eksekusi
*/

/* Kode Program Defer */

package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	// kita tidak peduli mau error/succsess kita tetap mau panggil function logging()
	// oleh karena itu kita menggunakan defer harus di awal function
	defer logging() // ini untuk memberitahu go-lang bahwa function runApplication() maka harus di akhiri dengan memanggil function logging()
	// jadi ketika memanggil panic: runtime error: integer divide by zero
	// dia tetap bisa memanggil function logging()
	fmt.Println("Run Application")
	/* Kalau hanya 3 perintah dibawah ini apabila terjadi error
	maka function dibawahnya tidak dapat dijalankan alias
	panic: runtime error: integer divide by zero
	jadi bilangan 10 tidak dapat dibagi dengan 0*/
	result := 10 / value
	fmt.Println("Result", result)
	// logging()
}

// jadi ketika isi dari function runApplication() telah dijalankan maka program akan menjalankan function logging()
func main() {
	runApplication(3)
}
