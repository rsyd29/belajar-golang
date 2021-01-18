package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// ketika menemukan genap, kodenya akan continye dilanjutkan
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
