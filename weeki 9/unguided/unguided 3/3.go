package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, N int, bMin, bMax *float64) {

	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < N; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, N int) float64 {

	var total float64 = 0
	for i := 0; i < N; i++ {
		total += arrBerat[i]
	}
	return total / float64(N)
}

func main() {
	var N int
	var berat arrBalita
	var minBerat, maxBerat float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, N, &minBerat, &maxBerat)

	rata := rerata(berat, N)

	fmt.Printf("Berat balita minimum: %.2f kg\n", minBerat)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", maxBerat)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
