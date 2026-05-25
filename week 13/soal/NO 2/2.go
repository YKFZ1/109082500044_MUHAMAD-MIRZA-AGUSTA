package main

import "fmt"

type Pemain struct {
	Nama   string
	Gol    int
	Assist int
}

func main() {
	var n int
	fmt.Println("Masukkan Data Input :")
	if _, err := fmt.Scanln(&n); err != nil {
		return
	}

	daftarPemain := make([]Pemain, n)
	for i := 0; i < n; i++ {
		var p1, p2, p3 string
		fmt.Scanln(&p1, &p2, &p3)

		var nama string
		var gol, assist int

		if p3 == "" {
			nama = p1
			fmt.Sscanf(p2, "%d", &gol)
			var temp string
			fmt.Scanln(&temp)
			fmt.Sscanf(temp, "%d", &assist)
		} else {
			var g, a int
			nGol, _ := fmt.Sscanf(p2, "%d", &g)
			nAst, _ := fmt.Sscanf(p3, "%d", &a)

			if nGol > 0 && nAst > 0 {
				nama = p1
				gol = g
				assist = a
			} else {
				nama = p1 + " " + p2
				fmt.Sscanf(p3, "%d", &gol)
				fmt.Scanln(&assist)
			}
		}

		daftarPemain[i] = Pemain{Nama: nama, Gol: gol, Assist: assist}
	}

	for i := 1; i < len(daftarPemain); i++ {
		key := daftarPemain[i]
		j := i - 1
		for j >= 0 {
			tukar := false
			if key.Gol > daftarPemain[j].Gol {
				tukar = true
			} else if key.Gol == daftarPemain[j].Gol && key.Assist > daftarPemain[j].Assist {
				tukar = true
			}

			if tukar {
				daftarPemain[j+1] = daftarPemain[j]
				j--
			} else {
				break
			}
		}
		daftarPemain[j+1] = key
	}

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %d\n", daftarPemain[i].Nama, daftarPemain[i].Gol, daftarPemain[i].Assist)
	}
}
