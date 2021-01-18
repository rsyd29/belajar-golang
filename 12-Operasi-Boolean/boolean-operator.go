/*
Operasi Boolean

Berikut adalah contoh code program boolen operator
1. && artinya dan untuk 2 boolean atau lebih
2. || artinya atau untuk 2 boolean atau lebih
3. !  artinya kebalikan untuk 1 boolean

Operasi && (dan) nilai kiri dan kanannya benar maka benar, kalau salah satunya ada salah
maka akan salah

Operasi || (atau) nilai kiri dan kanannya salah maka salah, kalau salah satunya ada benar
maka akan benar

Operasi ! (kebalikan) kalau nilai awalnya benar maka kebalikannya salah
*/

/*
Operasi && (jika semuanya true maka hasilnya true)
Nilai1 --> Operator --> Nilai2 --> Hasil
true	-->	&&	-->	true 	--> true
true	--> &&	--> false 	--> false
false	--> && 	--> true	--> false
false	--> && 	--> false	--> false
*/

/*
Operasi || (jika salah satu true maka hasilnya true)
Nilai1 --> Operator --> Nilai2 --> Hasil
true	-->	||	-->	true 	--> true
true	--> ||	--> false 	--> true
false	--> || 	--> true	--> true
false	--> || 	--> false	--> false
*/

/*
Operasi ! (jika salah satu true maka hasilnya true)
Operator --> Nilai2 --> Hasil
!	-->	true 	--> false
! 	--> false 	--> true

*/
package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 75

	var lulusUjian = ujian >= 80
	fmt.Println(lulusUjian)
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	//Biasanya pakai cara seperti ini agar lebih cepat
	fmt.Println(ujian >= 80 && absensi >= 80)
}
