package main

import "fmt"

type Mahasiswa struct {
	NIM   int
	Nama  string
	Nilai int
}

type ArrayMahasiswa struct {
	Data [51]Mahasiswa
	N    int
}

func main() {
	var arr ArrayMahasiswa
	var cariNIM int

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&arr.N)

	if arr.N > 51 {
		fmt.Println("Jumlah data melebihi kapasitas maksimal (51)")
		return
	}

	for i := 0; i < arr.N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&arr.Data[i].NIM, &arr.Data[i].Nama, &arr.Data[i].Nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&cariNIM)

	nilaiPertama := cariNilaiPertama(arr, cariNIM)
	nilaiTerbesar := cariNilaiTerbesar(arr, cariNIM)

	if nilaiPertama == -1 && nilaiTerbesar == -1 {
		fmt.Printf("Mahasiswa dengan NIM %d tidak ditemukan\n", cariNIM)
	} else {
		fmt.Printf("Nilai pertama dari NIM %d adalah %d\n", cariNIM, nilaiPertama)
		fmt.Printf("Nilai terbesar dari NIM %d adalah %d\n", cariNIM, nilaiTerbesar)
	}
}

func cariNilaiPertama(arr ArrayMahasiswa, nim int) int {
	for i := 0; i < arr.N; i++ {
		if arr.Data[i].NIM == nim {
			return arr.Data[i].Nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(arr ArrayMahasiswa, nim int) int {
	nilaiMax := -1
	ditemukan := false

	for i := 0; i < arr.N; i++ {
		if arr.Data[i].NIM == nim {
			if !ditemukan {
				nilaiMax = arr.Data[i].Nilai
				ditemukan = true
			} else if arr.Data[i].Nilai > nilaiMax {
				nilaiMax = arr.Data[i].Nilai
			}
		}
	}

	if ditemukan {
		return nilaiMax
	}
	return -1
}
