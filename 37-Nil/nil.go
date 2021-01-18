package main

import "fmt"

/**
Nil
1. Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi
maka secara otomatis nilainya adalah null atau nil
2. Berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe data
tertentu, maka secara otomatis akan dibuatkan default value nya
3. Namun di Go-Lang ada data nil, yaitu data kosong
4. Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface,
function, map, slice, pointer dan channel
*/

// function dengan nama NewMap memiliki parameter name tipe data string, dan return valuenya berupa map yang memiliki key string dan valuenya string
func NewMap(name string) map[string]string {
	if name == "" { // apabila nilai kosong
		return nil // maka return valuenya nil
	} else { // jika kondisi false
		// maka return valuenya berupa map
		return map[string]string{
			// yang memiliki key name, dan valuenya berupa parameter name
			"name": name,
		}
	}
}

func main() {
	human := NewMap("Budiman")
	var person map[string]string = nil
	fmt.Println(person) // isinya map kosong yang tidak ada isinya
	fmt.Println(human)  // isinya key name dengan value Budiman

	// Buat apa sih nil? digunakan untuk pengecekan seperti contoh ini
	data := NewMap("")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}

	// kalau tanpa nil, kita akan melakukan pengecekan seperti ini
	var newPerson map[string]string
	if newPerson["name"] == "" { // hasilnya Data Kosong, tapi untuk melakukan pengecekannya seperti ini
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

	// kalau pakai nil, kita akan melakukan pengecekan seperti ini
	var newPerson1WithNil map[string]string = NewMap("Budiman")
	if newPerson1WithNil == nil { // ini lebih simple
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(newPerson1WithNil)
	}
}
