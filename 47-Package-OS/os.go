/**
Package OS
1. Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
2. Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen
(bisa digunakan di semua sistem operasi)
3. dokumentasinya adalah https://golang.org/pkg/os/
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // untuk mengambil data arguments, isinya adalah slice of string
	// berisikan lokasi data file yang telah di compile
	fmt.Println("Data Arguments:")
	fmt.Println(args[0]) // ini merupakan lokasi data file yang telah di compile
	// jalankan pada terminal go run os.go Budiman Rasyid Zainuddin
	// untuk membedakannya
	// hati hati ketika print harus di lihat ada berapa data slice nya jangan sampai lebih indexnya
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	// function Hostname() untuk mengambil data host nya, nama host dari sistem operasi yang dipakai
	hostname, err := os.Hostname() // return valuenya adalah name string, dan err error
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error:", err.Error())
	}

	// Getenv itu berguna untuk mensetting configuration si aplikasinya
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)

}
