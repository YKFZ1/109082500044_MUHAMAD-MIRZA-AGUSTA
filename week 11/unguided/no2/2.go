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

	firstCandidate := 0
	firstVotes := -1
	secondCandidate := 0
	secondVotes := -1

	for i := 1; i <= 20; i++ {
		if votes[i] > 0 {
			if votes[i] > firstVotes {

				secondVotes = firstVotes
				secondCandidate = firstCandidate

				firstVotes = votes[i]
				firstCandidate = i
			} else if votes[i] > secondVotes {

				secondVotes = votes[i]
				secondCandidate = i
			} else if votes[i] == firstVotes {

				if i < firstCandidate {

					secondVotes = firstVotes
					secondCandidate = firstCandidate

					firstCandidate = i
				} else if i < secondCandidate || secondCandidate == 0 {

					secondCandidate = i
				}
			} else if votes[i] == secondVotes {

				if i < secondCandidate {
					secondCandidate = i
				}
			}
		}
	}

	fmt.Printf("Ketua RT: %d\n", firstCandidate)
	fmt.Printf("Wakil ketua: %d\n", secondCandidate)
}
