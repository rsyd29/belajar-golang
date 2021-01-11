/*
	Function Return Value
	1. Function bisa mengembalikan data
	2. Untuk memberitahu bahwa function mengembalikan data,
	kita harus menuliskan tipe dat akembalian dari function
	tersebut
	3. Jika function tersebut kita deklarasikan dengan tipe data
	pengembalian, maka wajib di dalam function nya kita harus
	mengembalikan data
	4. Untuk mengembalikan data dari function, kita bisa
	menggunakan kata kunci return, diikuti dengan datanya
*/

package main

import "fmt"

func getHello(name string) string {
	// ini akan mengembalikan data yang tipe datanya string
	// biasanya return di simpan pada bagian bahwa, setelah return kode tersebut akan kembali lagi ke functionnya
	// Jadi kode blok di bawah return tidak akan pernah dieksekusi lagi
	return "Hello " + name
}

/*
 Kita juga bisa menggunakan return di dalam kondisi
*/
func getHelloBaru(name string) string {
	if name == "Budi" {
		return "Hello Budi" // return di dalam kondisi
	}
	return "Hello " + name

}

func main() {
	// ketika kita memanggil func getHello maka itu akan mengembalikan data
	// oleh karena itu kita bisa menyimpan datanya ke dalam variable
	result := getHello("Budiman")
	resultBaru := getHelloBaru("Rasyid")
	fmt.Println(result)
	fmt.Println(resultBaru)
}
