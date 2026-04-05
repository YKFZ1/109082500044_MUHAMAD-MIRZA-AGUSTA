package main

import "fmt"

func tampilkanGanjil(current, n int) {

	if current > n {
		return
	}

	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}

	tampilkanGanjil(current+1, n)
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Masukkan bilangan bulat positif!")
		return
	}

	fmt.Printf("Bilangan ganjil dari 1 hingga %d adalah: ", n)
	tampilkanGanjil(1, n)
	fmt.Println()
}
