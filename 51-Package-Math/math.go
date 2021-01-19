/**
Package Math (Matematika)
1. Package math merupakan package yang berisikan constant dan fungsi matematika
2. Dokumentasinya adalah https://golang.org/pkg/math/

Beberapa Function di Package math
Function --> Kegunaan
math.Round(float64) --> Membulatkan float64 ke atas atau ke bawah, sesuai dengan yang paling dekat
math.Floor(float64) --> Membulatkan float64 ke bawah (di paksa)
math.Ceil(float64) --> Membulatkan float64 ke atas
math.Max(float64, float64) --> Mengembalikan nilai float64 paling besar
math.Min(float64, float64) --> Mengembalikan nilai float64 paling kecil
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7)) // 2 dibulatkan ke atas
	fmt.Println(math.Round(1.3)) // 1 dibulatkan ke bawah

	fmt.Println(math.Floor(1.7)) // 1 tetap dibulatkan ke bawab
	fmt.Println(math.Ceil(1.3))  // 2 tetap dibulatkan ke atas

	fmt.Println(math.Max(10, 20)) // 20 mencari paling besar
	fmt.Println(math.Min(10, 20)) // 10 mencari paling kecil
}
