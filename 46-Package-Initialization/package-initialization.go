package main

import _ "belajar-golang/46-Package-Initialization/database"

// setelah import tambahkan _ agar tidak terjadi error
// error karena package database tidak pernah terpakai
// oleh karena itu tambahkan _ (underscore) setelah import
// istilahnya Blank Identifier
// apabila ingin menggunakan package database tersebut maka hapus _ underscore nya

func main() {
	//result:= database.GetDatabase()
	//fmt.Println(result) // outpunya MySQL, padahal function init() blm pernah dieksekusi pada file database.go
}
