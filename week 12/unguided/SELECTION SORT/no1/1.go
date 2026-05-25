package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

func main() {
	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var m int

		fmt.Scan(&m)

		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(rumah)

		for k := 0; k < m; k++ {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[k])
		}
		fmt.Println()
	}
}
