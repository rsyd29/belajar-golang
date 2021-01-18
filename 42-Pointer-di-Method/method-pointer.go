/**
Pointer di Method
1. Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses
di dalam method adalah pass by value
2. Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory
karena harus selalu diduplikasi ketika memanggil method
*/

package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() { // ini struct function/method, dan ini merupakan pass by value
	man.Name = "Mr. " + man.Name
	// artinya
	fmt.Println("Di Method tanpa pointer", man.Name)
}

//biasanya orang orang ketika menggunakan struct function/method menggunakan pointer
func (man *Man) MarriedPointer() { // ini struct function/method, dan ini merupakan struct method dengan pointer
	man.Name = "Mr. " + man.Name
	// artinya
	fmt.Println("Di Method dengan pointer", man.Name)
}

func main() {
	budiman := Man{"Budiman"} // ini tidak berubah valuenya
	// data budiman di duplikat lalu dikirim ke function Married()
	budiman.Married()
	fmt.Println("tanpa pointer:", budiman) // hasilnya tetap Budiman

	rasyid := Man{"Rasyid"} // ini akan berubah valuenya
	// karena menggunakan pointer
	rasyid.MarriedPointer()
	fmt.Println("dengan pointer:", rasyid)

	/**
	Ingat ketika membuat struct method atau function usahakan menggunakan pointer
	dibandingkan dengan value biasa, selain itu lebih hemat memory tidak menduplikat
	datanya
	*/
}
