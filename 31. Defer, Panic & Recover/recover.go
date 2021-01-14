/*
	Recover
	1. Recover adalah function yang bisa kita gunakan untuk
	menangkap data panic
	2. Dengan recover proses panic akan terhenti, sehingga
	program akan tetap berjalan
*/

package main

import "fmt"

func endApp() {
	// recover masukkan ke dalam defer function, kalau tidak programnya keburu panic jadi terdapat error
	message := recover() // function recover() digunakan untuk menangkap data dari function panic()
	// function recover() juga digunakan agar program tetap berjalan walaupun ada function panic() atau error
	if message != nil { // apabila terdapat message recover yang ada pada panic itu tidak sama dengan nil maka
		// menampilan pesan error yang ada pada function panic()
		fmt.Println("Error dengan message:", message)
	}
	// jika tidak maka aplikasi berjalan
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp() // perintah ini diabaikan terlebih dahulu, sebelum perintah bahwanya selesai, jika sudah selesai maka akan di jalankan
	if error {
		// apabila error bernilai true
		panic("APLIKASI ERROR") // ketika berhenti maka diakan kembali ke function runApp()
		// oleh karena itu kita buat function recover di luar function ini yaitu function endApp()
	}
	// apabila error bernilai false
	fmt.Println("Aplikasi Berjalan")
}

/*
panic ini berfungsi untuk apabila aplikasi yang dibuat
error lalu ingin aplikasi dibuat berhenti
*/
func main() {
	runApp(false)                           // jika error true maka aplikasi selesai dan menampilkan panic bahwa APP ERROR
	fmt.Println("Budiman Rasyid Zainuddin") // untuk cek apakah program berhenti atau tidak
}
