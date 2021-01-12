/*
	Variadic Function
	1. Parameter yang berada di posisi terakhir, memiliki
	kemampuan dijadikan sebuah varargs (variable arguments)
	2. Varargs artinya datanya bisa menerima lebih dari satu
	input, atau anggap saja semacam Array
	(seperti dynamic)
	3. Apa bedanya dengan parameter biasa dengan tipe data
	Array?
		- Jika Parameter tipe Array, kita wajib membuat Array
		terlebih dahulu sebelum mengirimkan ke function
		- Jika parameter menggunakan varargs, kita bisa langsung
		mengirim datanya, jika lebih dari satu, cukup gunakan
		tanda koma
*/

package main

import "fmt"

// kalau ingin membuat lebih dari 1 parameter ditambah dengan variable argument maka variable argument harus ditempatkan di paling akhir
func sumAll(numbers ...int) int { // ini sebuah variadic function
	// kenapa disebut variadic function karena parameternya itu memilih variable arguments
	// cirinya ada titik tiga (...)
	// artinya kita dapat memanggil numbers sebanyak lebih dari sekali
	// setelah itu parameter dari variable function tersebut menjadi sebuah slice
	total := 0
	for _, number := range numbers { // ini merupakan iterasi dari range numbers yaitu slice numbers tersebut
		// index diganti menjadi _ karena tidak terpakai kita ignore saja
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 90, 50, 40, 10) // kalau lebih dari satu tinggal tambahkan koma (,)
	fmt.Println(total)

	/*
		Slice Parameter
		1. Kadang ada kasus dimana kita menggunakan variadic function,
		namun memiliki variable berupa slice
		2. Kita bisa menjadikan slice sebagai vararg (variable arguments)
		parameter
	*/

	slice := []int{10, 10, 10, 10, 10, 10} // ini merupakan slice
	total = sumAll(slice...)               // ini untuk memecahkan sebuah slice menjadi tipe variable arguments, karena function sumAll itu parameternya variable arguments
	fmt.Println(total)
}
