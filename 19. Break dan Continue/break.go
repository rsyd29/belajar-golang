/*
	Break & Continue
	1. Break & continue adalah kata kunci yang bisa digunakan
	dalam perulangan
	2. Break digunakan untuk menghentikan seluruh perulangan
	3. Continue digunakan untuk menghentikan perulangan yang
	berhalan, dan langsung melanjutkan ke perulangan selanjutnya
*/

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
			// apabila kondisi terpenuhi maka break akan menghentikan perulangannya
		}
		fmt.Println("Perulangan ke", i)
	}
}
