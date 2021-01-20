/**
Package Sort
1. Package sort adalah package yang berisikan utilitas untuk proses pengurutan
2. Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface
sort.Interface
3. Dokumentasinya adalah https://golang.org/pkg/sort/
*/

package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User // membuat alias menjadi Slice

func (value UserSlice) Len() int {
	return len(value) // ini mengukur panjang dari slice
}

func (value UserSlice) Less(i, j int) bool { // ini akan dibandingkan antara index i dan j
	return value[i].Age < value[j].Age // bandingkan berdasarkan Age
}

func (value UserSlice) Swap(i, j int) { // ini akan merubah tempatnya
	value[i], value[j] = value[j], value[i] // cara cepatnya
}

func main() {
	users := []User{
		{"Budiman", 20},
		{"Zainuddin", 53},
		{"Nuratitin", 50},
		{"Subhan Amin", 25},
	}

	fmt.Println(users)          // sebelum proses sorting
	sort.Sort(UserSlice(users)) // ini dikonversi menjadi alias
	fmt.Println(users)          // setelah proses sorting
}
