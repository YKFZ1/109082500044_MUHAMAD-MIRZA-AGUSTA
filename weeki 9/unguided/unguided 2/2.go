package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var berat [1000]float64

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	var totalPerWadah [1000]float64

	idxIkan := 0
	for w := 0; w < jumlahWadah; w++ {
		for i := 0; i < y && idxIkan < x; i++ {
			totalPerWadah[w] += berat[idxIkan]
			idxIkan++
		}
	}

	for w := 0; w < jumlahWadah; w++ {
		fmt.Printf("%.2f", totalPerWadah[w])
		if w < jumlahWadah-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	var totalSemua float64
	for w := 0; w < jumlahWadah; w++ {
		totalSemua += totalPerWadah[w]
	}
	rataRata := totalSemua / float64(jumlahWadah)
	fmt.Printf("%.2f\n", rataRata)
}
