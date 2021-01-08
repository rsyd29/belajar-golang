package main

import "fmt"

/*
- () di dalam kurung buka dan kurung tutup itu merupakan parameter dari function
- parameter ini bersifat tidaklah wajib, boleh tidak menggunakannya.
- ketika membuat function itu kadang2 membutuhkan data dari luar, maka dibuatlah parameter
- untuk menamakan parameter itu mirip seperti menamakan variabel, 
- jadi terdapat nama parameternya, lalu diikuti dengan tipe datanya
-
kalau lebih dari satu parameter, silahkan tambahkan koma
- nama parameter unik, apabila ada 2 parameter atau lebih maka nama parameter tidak boleh sama, harus berbeda.
- urutannya harus sama ketika ingin memanggil parameternya pada nama parameter.
*/

func sayHelloTo(firstName string, lastName string){ 
	// func yang memiliki 2 parameter
	fmt.Println("Hello", firstName, lastName)
}


func main(){
	firstName := "Budiman" // boleh dimasukan dgn menggunakan variable
	sayHelloTo(firstName, "Rasyid") // urutan nama parameter harus sesuai
	// dan tipe data yang sesuai

	sayHelloTo("Nuratitin", "Suwidi")
	sayHelloTo("Zainuddin", "Fathullah")
	sayHelloTo("Subhan", "Amin")
}
