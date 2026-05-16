package main

import "fmt"

func main() {
	var input int
	votes := make(map[int]int)
	validCount := 0
	totalCount := 0

	for {
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		totalCount++

		if input >= 1 && input <= 20 {
			validCount++
			votes[input]++
		}

	}

	fmt.Printf("Suara masuk: %d\n", totalCount)
	fmt.Printf("Suara sah: %d\n", validCount)

	for i := 1; i <= 20; i++ {
		if votes[i] > 0 {
			fmt.Printf("%d: %d\n", i, votes[i])
		}
	}
}
