/*
	Panic
	1. Panic function adalah function yang bisa kita gunakan
	untuk menghentikan program
	2. Panic function biasanya dipanggil ketika terjadi error
	pada saat program kita berjalan
	3. Saat panic function dipanggil, program akan terhenti,
	namun defer function tetap akan dieksekusi
*/

package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp() // perintah ini diabaikan terlebih dahulu, sebelum perintah bahwanya selesai, jika sudah selesai maka akan di jalankan
	if error {
		// apabila error bernilai true
		panic("APLIKASI ERROR")
	}
	// apabila error bernilai false
	fmt.Println("Aplikasi Berjalan")
}

/*
panic ini berfungsi untuk apabila aplikasi yang dibuat
error lalu ingin aplikasi dibuat berhenti
*/
func main() {
	runApp(true) // jika error true maka aplikasi selesai dan menampilkan panic bahwa APP ERROR
}
