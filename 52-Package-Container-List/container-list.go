/**
Package Container List
1. Package container/list adalah implementasi struktur data double linked list di Go-Lang
2. Dokumentasinya adalah https://golang.org/pkg/container/list/
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// di list tidak ada index
	data := list.New()       // untuk membuat list
	data.PushBack("Budiman") // untuk membuat data baru
	data.PushBack("Rasyid")
	data.PushBack("Zainuddin")
	data.PushFront("Mr.") // untuk membuat data baru dan di masukkan ke paling depan

	// ambil data pertama
	//fmt.Println(data.Front().Value)
	// ambil data terakhir
	//fmt.Println(data.Back().Value)

	// untuk melihat next data nya
	//data.Front().Next().Next() // perhatikan ketika Next() suatu saat akan nil datanya

	// untuk melihat prev data nya
	//data.Front().Prev().Prev() // perhatikan ketika Prev() suatu saat akan nil datanya

	// contoh kasusnya
	//fmt.Println(data.Front().Prev()) // akan terjadi nil
	//fmt.Println(data.Back().Next()) // akan terjadi nil

	// untuk mengiterasi kan data di list dari depan
	for elementFront := data.Front(); elementFront != nil; elementFront = elementFront.Next() {
		fmt.Println(elementFront.Value)
	}

	// untuk mengiterasi kan data di list dari belakang
	for elementBack := data.Back(); elementBack != nil; elementBack = elementBack.Prev() {
		fmt.Println(elementBack.Value)
	}
}
