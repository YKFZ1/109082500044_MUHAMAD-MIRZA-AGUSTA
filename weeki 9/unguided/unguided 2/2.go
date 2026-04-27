package main

import (
	"fmt"
	"math"
)

func main() {
	var N int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	arr := make([]int, N)

	// Mengisi array
	fmt.Println("Masukkan", N, "bilangan bulat:")
	for i := 0; i < N; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// Menu pilihan
	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Tampilkan seluruh isi array")
		fmt.Println("2. Tampilkan elemen indeks ganjil")
		fmt.Println("3. Tampilkan elemen indeks genap")
		fmt.Println("4. Tampilkan elemen indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata")
		fmt.Println("7. Tampilkan standar deviasi")
		fmt.Println("8. Tampilkan frekuensi suatu bilangan")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu (1-9): ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:

			fmt.Print("Seluruh isi array: ")
			for i := 0; i < N; i++ {
				fmt.Printf("%d ", arr[i])
			}
			fmt.Println()

		case 2:

			fmt.Print("Elemen dengan indeks ganjil: ")
			ada := false
			for i := 1; i < N; i += 2 {
				fmt.Printf("%d ", arr[i])
				ada = true
			}
			if !ada {
				fmt.Print("(tidak ada)")
			}
			fmt.Println()

		case 3:

			fmt.Print("Elemen dengan indeks genap: ")
			for i := 0; i < N; i += 2 {
				fmt.Printf("%d ", arr[i])
			}
			fmt.Println()

		case 4:

			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			if x <= 0 {
				fmt.Println("x harus lebih besar dari 0")
				break
			}
			fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
			ada := false
			for i := x; i < N; i += x {
				fmt.Printf("%d ", arr[i])
				ada = true
			}
			if !ada {
				fmt.Print("(tidak ada)")
			}
			fmt.Println()

		case 5:

			var idx int
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scan(&idx)
			if idx >= 0 && idx < N {

				arr = append(arr[:idx], arr[idx+1:]...)
				N--
				fmt.Print("Isi array setelah dihapus: ")
				for i := 0; i < N; i++ {
					fmt.Printf("%d ", arr[i])
				}
				fmt.Println()
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case 6:

			if N == 0 {
				fmt.Println("Array kosong, tidak bisa menghitung rata-rata")
				break
			}
			sum := 0
			for i := 0; i < N; i++ {
				sum += arr[i]
			}
			rata := float64(sum) / float64(N)
			fmt.Printf("Rata-rata: %.2f\n", rata)

		case 7:

			if N == 0 {
				fmt.Println("Array kosong, tidak bisa menghitung standar deviasi")
				break
			}

			sum := 0
			for i := 0; i < N; i++ {
				sum += arr[i]
			}
			rata := float64(sum) / float64(N)

			var variance float64
			for i := 0; i < N; i++ {
				variance += math.Pow(float64(arr[i])-rata, 2)
			}
			variance /= float64(N)
			stdDev := math.Sqrt(variance)
			fmt.Printf("Standar deviasi: %.2f\n", stdDev)

		case 8:

			if N == 0 {
				fmt.Println("Array kosong")
				break
			}
			var cari int
			fmt.Print("Masukkan bilangan yang dicari frekuensinya: ")
			fmt.Scan(&cari)
			frekuensi := 0
			for i := 0; i < N; i++ {
				if arr[i] == cari {
					frekuensi++
				}
			}
			fmt.Printf("Frekuensi bilangan %d di dalam array: %d kali\n", cari, frekuensi)

		case 9:
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
