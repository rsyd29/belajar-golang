package main

import "belajar-golang/44-Package-dan-Import/helper"
import "belajar-golang/44-Package-dan-Import/other"

// nama foldernya harus sesuaikan dengan gopathnya
// ~/go/src/belajar-golang

func main() {
	helper.SayHello("Budiman")
	other.SayHello("Rasyid")
}
