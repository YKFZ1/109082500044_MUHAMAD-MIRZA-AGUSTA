package main

import "fmt"

func tampilkanFaktor(n, current int) {

	if current > n {
		return
	}

	if n%current == 0 {
		fmt.Printf("%d ", current)
	}

	tampilkanFaktor(n, current+1)
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Masukkan bilangan bulat positif!")
		return
	}

	fmt.Printf("Faktor dari %d adalah: ", n)
	tampilkanFaktor(n, 1)
	fmt.Println()
}
