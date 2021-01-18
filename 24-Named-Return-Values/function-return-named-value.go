/*
	Named Return Values (Hanya di Golang)
	1. Biasanya saat kita memberi tahu bahwa sebuah function
	mengembalikan value, maka kita hanya mendeklarasikan tipe
	data return value di function
	2. Namun kita juga bisa membuat variable secara langsung
	di tipe data return function nya
*/

package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	// kita bisa langsung membuat var di parameternya seperti di atas
	firstName = "Budiman"
	middleName = "Rasyid"
	lastName = "Zainuddin"

	// return firstName, middleName, lastName
	// kita juga bisa langsung ketikkan
	return
	// tanpa harus memasukkan semua var
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
