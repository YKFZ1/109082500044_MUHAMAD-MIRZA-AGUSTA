package main

import "fmt"

func insertionSort(data []int) {
	/* I.S. terdefinisi slice data yang berisi n bilangan bulat
	   F.S. slice data terurut secara mengecil atau descending dengan INSERTION SORT */

	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i

		for j > 0 && temp > data[j-1] {
			data[j] = data[j-1]
			j--
		}

		data[j] = temp
	}
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		return
	}

	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	insertionSort(data)

	for i, v := range data {
		fmt.Print(v)
		if i < n-1 {
			fmt.Print(" ")
		}
	}
}
