package main

import "fmt"

func main() {
	var target, tabungan, total, hari int
	fmt.Print("Masukkan target tabungan: ")
	fmt.Scan(&target)

	total = 0
	hari = 0
	for total < target {
		hari++
		fmt.Print("Masukkan jumlah tabungan hari ke: ", hari)
		fmt.Scan(&tabungan)

		total = total + tabungan
	}
	fmt.Printf("Selamat! Anda telah mencapai target tabungan dalam %d hari.\n", hari)
	fmt.Printf("Total tabungan terkumpul: %d\n", total)
}
