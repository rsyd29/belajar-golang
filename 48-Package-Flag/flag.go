/**
Package Flag terkait erat dengan package os
1. Package flag berisikan fungsionalitas untuk memparsing command line argument
2. dokumentasinya adalah https://golang.org/pkg/flag/
*/

package main

import (
	"flag"
	"fmt"
)

func main() {
	// untuk menjalankannya
	// go run flag.go -host=localhost -username=root -password=root
	// lalu bagaimana agar bisa otomatis? dengan melakukan perintah di bawah ini

	/**
	flag.String() ada 3 parameter, jadi ini akan di parsing ke dalam string
	1. name --> namanya
	2. value --> isi dari namanya
	3. usage --> helper untuk penggunaannya
	ini hasilnya akan menjadi pointer string (*string)
	*/
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your number")

	// ini untuk melakukan proses parsingnya sebelum menggunakan variable host, username, dan password
	flag.Parse() // sebelum gunakan flag pastikan harus gunakan flag.Parse()

	// gunakan variablenya, hasil 3 variable tersebut berupa pointer, jadi harus menggunakan tanda *
	fmt.Println("Host:", *host)
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Number:", *number)

	// ini berguna untuk membuat aplikasi yang berbasis terminal
}
