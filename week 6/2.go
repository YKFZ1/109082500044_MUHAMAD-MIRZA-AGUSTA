package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {

	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
}

func biayaPerHariDomestik(jumlahMhs int) int {

	biayaMakan := 35000 * 2
	biayaPenginapan := 250000
	uangSaku := 300000

	biayaPerMhsPerHari := biayaMakan + biayaPenginapan + uangSaku

	return biayaPerMhsPerHari * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {

	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)

	biayaDomestikPerHari := float64(biayaPerHariDomestik(jumlahMhs))

	if tujuan == "domestik" {
		*totalBiaya = biayaDomestikPerHari * float64(hariDitanggung)
	} else {
		*totalBiaya = biayaDomestikPerHari * 1.5 * float64(hariDitanggung)
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}
