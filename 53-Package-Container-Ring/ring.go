/**
Package Container/ring
1. Package container/ring adalah implementasi struktur data circular list
2. Circular list adalah struktur data ring, dimana diakhiri element akan kembali ke
element awal (HEAD)
3. Dokumentasinya adalah https://golang.org/pkg/container/ring/
*/

package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// manualnya seperti ini
	var oldData *ring.Ring = ring.New(5)
	oldData.Value = "Budiman"

	var oldData2 = oldData.Next()
	oldData2.Value = "Rasyid"

	// untuk melakukan otomatis iterasi seperti ini
	data := ring.New(5) // untuk membuat ring baru, ditentukan rangenya
	// var data *ring.Ring = ring.New(5) // ini sama seperti code di atas
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10) // untuk mengambil datanya dan di konversi dari int ke string
		data = data.Next()                                     // untuk mengambil data selanjutnya
	}

	// jangan pakai data for loop, nanti tidak akan berhenti karena ada data.Next()

	data.Do(func(value interface{}) { // function bawaan ring, untuk melakukan iterasi
		fmt.Println(value)
	})
}
