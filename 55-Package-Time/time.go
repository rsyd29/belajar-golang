/**
Package Time
1. Package Time adalah package yang berisikan fungsionalitas untuk management waktu di
Go-Lang
2. Dokumentasinya adalah https://golang.org/pkg/time/

Beberapa Function di Package Time
Function --> Kegunaan
time.Now() --> Untuk mendapatkan waktu saat ini
time.Date(...) --> Untuk membuat waktu secara manual
time.Parse(layout, string) --> Untuk memparsing waktu dari string
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)              // semua format time
	fmt.Println(now.Year())       // tahun
	fmt.Println(now.Month())      // bulan
	fmt.Println(now.Day())        // hari
	fmt.Println(now.Hour())       // jam
	fmt.Println(now.Minute())     // menit
	fmt.Println(now.Second())     // detik
	fmt.Println(now.Nanosecond()) // mili detik

	// untuk membuat waktu sesuai keinginan
	utc := time.Date(2021, time.November, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	// perhatikan ketika membuat layout
	layout := "2006-01-02" // layoutnya pakai tanggal juga, bukan pakai format yyyy/mm/dd
	parse, _ := time.Parse(layout, "2021-01-20")
	fmt.Println(parse)
	newParse, _ := time.Parse(layout, "2000-03-29")
	fmt.Println(newParse)
}
