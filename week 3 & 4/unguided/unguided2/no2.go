package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai z : ")
	fmt.Scan(&c)

	result1 := f(g(h(a)))
	fmt.Printf("F(G(H( %d ))) : %d\n", a, result1)

	result2 := g(h(f(b)))
	fmt.Printf("G(H(F( %d ))) : %d\n", b, result2)

	result3 := h(f(g(c)))
	fmt.Printf("H(F(G( %d ))) : %d\n", c, result3)
}
