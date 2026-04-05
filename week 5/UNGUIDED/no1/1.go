package main

import "fmt"

func fibonacci(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int

	fmt.Print("Masukkan suku ke-n yang ingin dicari: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Masukkan bilangan non-negatif!")
		return
	}

	fmt.Printf("\nNilai suku ke-%d dari deret Fibonacci adalah: %d\n\n", n, fibonacci(n))

	fmt.Println("Deret Fibonacci hingga suku ke-n:")

	fmt.Print("n     |")
	for i := 0; i <= n; i++ {
		fmt.Printf(" %d |", i)
	}
	fmt.Println()

	fmt.Print("-------")
	for i := 0; i <= n; i++ {
		fmt.Print("----")
	}
	fmt.Println()

	fmt.Print("S_n   |")
	for i := 0; i <= n; i++ {
		fmt.Printf(" %d |", fibonacci(i))
	}
	fmt.Println()
}
