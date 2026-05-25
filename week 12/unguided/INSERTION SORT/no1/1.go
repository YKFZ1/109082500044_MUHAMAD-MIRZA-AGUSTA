package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var arr []int

	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}

		if num < 0 {
			break
		}

		arr = append(arr, num)
	}

	if len(arr) == 0 {
		return
	}

	insertionSort(arr)

	for i, val := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()

	isConstant := true
	var delta int

	if len(arr) > 1 {

		delta = arr[1] - arr[0]
		for i := 1; i < len(arr)-1; i++ {
			if (arr[i+1] - arr[i]) != delta {
				isConstant = false
				break
			}
		}
	} else {

		delta = 0
	}

	if isConstant {
		fmt.Printf("Data berjarak %d\n", delta)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
