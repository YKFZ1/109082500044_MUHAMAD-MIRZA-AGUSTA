package main

import (
	"fmt"
)

// Selection sort untuk mengurutkan membesar (ascending)
func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Selection sort untuk mengurutkan mengecil (descending)
func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for daerah := 0; daerah < n; daerah++ {
		var m int
		fmt.Scan(&m)

		// Baca m bilangan
		rumah := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Scan(&rumah[i])
		}

		// Pisahkan nomor ganjil dan genap
		var ganjil, genap []int
		for _, num := range rumah {
			if num%2 == 1 {
				ganjil = append(ganjil, num)
			} else {
				genap = append(genap, num)
			}
		}

		// Urutkan ganjil membesar, genap mengecil
		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		// Cetak hasil: ganjil dulu, lalu genap
		for i, v := range ganjil {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		for i, v := range genap {
			if len(ganjil) > 0 || i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}
