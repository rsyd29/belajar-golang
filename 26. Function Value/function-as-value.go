/*
	Function as Value
	1. Function adalah first class citizen
	(tidak ada anak tiri di golang terhadap function, beda
	dengan Java function harus berada di dalam class,
	sedangkan di golang dia berdiri sendiri)
	2. Function juga merupakan tipe data, dan bisa disimpan di
	dalam variable
*/

package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye
	// menyimpan function getGoodBye ke dalam variable goodbye
	// jadi function getGoodBye disimpan sebagai value untuk variable goodbye
	result := goodbye("Budiman") // karena dia return value bertipe data string
	fmt.Println(result)
	// jadi goodbye bisa kasih parameter seperti getGoodBye
	// karena goodbye memiliki value function getGoodBye

	/*
		Keuntungannya ketika kita membuat sebuah function
		yang membutuhkan function lain, kita bisa melakukan
		hal seperti ini, karena function bisa di set sebagai
		value, lalu kita bisa masukkan ke dalam variable
	*/
}
