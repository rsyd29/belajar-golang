package main

import "fmt"

func sayHello(){
	fmt.Println("Hello World!")
}

func main(){
	for i:=0; i<10; i++{
		sayHello() 
		/* 
		ini merupakan func sayHello yang sedang dipanggil
		di func main, maka func main ini akan memanggil func sayHello
		sebanyak looping for tersebut
		*/
	}
}
