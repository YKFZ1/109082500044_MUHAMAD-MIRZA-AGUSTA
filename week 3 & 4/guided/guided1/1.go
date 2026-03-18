package main

import "fmt"

func hitungluaspersegi(panjang int, lebar int) int {
	luas := panjang * lebar
	return luas
}

func main() {
	p := 10
	l := 5

	hasil := hitungluaspersegi(p, l)
	fmt.Printf("Luas persegi dengan panjang %d dan lebar %d adalah: %d\n", p, l, hasil)
}
