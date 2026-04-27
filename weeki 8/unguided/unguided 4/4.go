package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {

	var kar rune
	*n = 0

	for *n < NMAX {
		fmt.Scanf("%c", &kar)
		if kar == '.' || kar == '\n' {
			break
		}
		t[*n] = kar
		*n++
	}
}

func cetakArray(t tabel, n int) {

	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {

	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrome(t tabel, n int) bool {

	var temp tabel
	for i := 0; i < n; i++ {
		temp[i] = t[i]
	}
	balikanArray(&temp, n)

	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Masukkan karakter (akhiri dengan titik): ")

	isiArray(&tab, &m)

	fmt.Print("Isi array: ")
	cetakArray(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Setelah dibalik: ")
	cetakArray(tab, m)

	if palindrome(tab, m) {
		fmt.Println("Merupakan PALINDROME")
	} else {
		fmt.Println("BUKAN palindrome")
	}
}
