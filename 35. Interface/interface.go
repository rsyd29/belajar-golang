/*
	Interface
	1. Interface adalah tipe data Abstract, dia tidak memiliki
	implementasi langsung
	2. Sebuah interface berisikan definisi-definisi method
	3. Biasanya interface digunakan sebagai kontrak
*/

package main

import "fmt"

type HasName interface { // ini merupakan interface dengan nama HasName
	// dia memiliki kontrak dengan sebuah method/function bernama GetName() return value tipe data string
	GetName() string // jadi kita harus membuat function ini dengan return value string
}

// interface itu sama seperti struct, jadi kita bisa memakainya sebagai parameter
func SayHello(hasName HasName) { // parameternya adalah interface HasName dengan nama hasName
	fmt.Println("Hello", hasName.GetName()) // ini kita panggil sebuah parameter dengan method/function GetName()
}

/*
	Implementasi Interface
	1. Setiap tipe data yang sesuai dengan kontrak interface,
	secara otomatis dianggap sebagai interface tersebut
	2. Sehingga kita tidak perlu mengimplementasikan interface
	secara manual
	3. Hal ini agak berbeda dengan bahasa pemrograman lain yang
	ketika membuat interface, kita harus menyebutkan secara
	eksplisit akan menggunakan interface mana
*/

type Person struct {
	Name string
}

// func ini untuk kontrak dari interface HasName
func (person Person) GetName() string {
	return person.Name
}

// kita juga bisa menggunakan struct yang berbeda
type Animal struct {
	Name string
}

// untuk memanggil function GetName() yang sama
// lalu beda structnya juga bisa digunakan
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Budiman Rasyid"}
	// parameter Function SayHello itu kontraknya adalah function GetName() pada interface HasName
	// artinya adalah tipe data apapun yang mempunyai function GetName() return value string
	// maka dia berhak mengikuti kontrak HasName
	SayHello(person)

	cat := Animal{
		Name: "Pushy",
	}
	SayHello(cat)
}
