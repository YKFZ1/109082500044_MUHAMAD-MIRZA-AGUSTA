package main

import "fmt"

type arrdata [5]string // tipe data alias

func seqsearch(arr arrdata, binatangcari string) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == binatangcari {
			return i
		}
	}
	return -1
}

func main() {
	var arrbinatang arrdata

	for i := 0; i < len(arrbinatang); i++ {
		fmt.Printf("Masukkan data binatang ke-%d : ", i)
		fmt.Scan(&arrbinatang[i])
	}

	fmt.Println()

	var binatangcari string
	fmt.Print("Masukkan nama binatang yang ingin dicari : ")
	fmt.Scan(&binatangcari)

	idxcari := seqsearch(arrbinatang, binatangcari)

	if idxcari != -1 {
		fmt.Printf("Binatang %s ditemukan pada index ke-%d\n", binatangcari, idxcari)
	} else {
		fmt.Printf("Binatang %s tidak ditemukan dalam array\n", binatangcari)
	}
}
	