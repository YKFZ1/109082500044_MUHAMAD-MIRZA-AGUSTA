package main

import "fmt"

type angka int

type kata string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var buku Buku
	var tempJudul, tempPenulis, tempPenerbit string
	var tempTahun, tempHalaman int

	fmt.Println("=== INPUT DATA BUKU ===")

	fmt.Print("Masukkan judul buku : ")
	fmt.Scanln(&tempJudul)

	fmt.Print("Masukkan nama penulis : ")
	fmt.Scanln(&tempPenulis)

	fmt.Print("Masukkan nama penerbit : ")
	fmt.Scanln(&tempPenerbit)

	fmt.Print("Masukkan tahun terbit: ")
	fmt.Scanln(&tempTahun)

	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scanln(&tempHalaman)

	buku.judul = kata(tempJudul)
	buku.penulis = kata(tempPenulis)
	buku.penerbit = kata(tempPenerbit)
	buku.tahunTerbit = angka(tempTahun)
	buku.jumlahHalaman = angka(tempHalaman)

	fmt.Println("\n=== INFORMASI BUKU ===")
	fmt.Println("Judul buku      :", buku.judul)
	fmt.Println("Penulis         :", buku.penulis)
	fmt.Println("Penerbit        :", buku.penerbit)
	fmt.Println("Tahun terbit    :", buku.tahunTerbit)
	fmt.Println("Jumlah halaman  :", buku.jumlahHalaman, "halaman")
}
