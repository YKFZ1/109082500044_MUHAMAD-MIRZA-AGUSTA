package main

import "fmt"

func main() {
	var beratGram float64

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&beratGram)

	kg := int(beratGram / 1000)
	sisaGram := beratGram - float64(kg)*1000

	biayaKg := float64(kg) * 10000

	var tambahanBiaya float64
	if sisaGram > 500 {

		tambahanBiaya = sisaGram * 5
	} else if sisaGram > 0 && sisaGram <= 500 {

		tambahanBiaya = sisaGram * 15
	}

	totalKg := beratGram / 1000
	if totalKg > 10 {
		tambahanBiaya = 0
	}

	totalBiaya := biayaKg + tambahanBiaya

	fmt.Println()
	fmt.Println("===== Detail Perhitungan =====")
	fmt.Printf("Detail berat : %d kg + %.0f gram\n", kg, sisaGram)
	fmt.Printf("Detail biaya : Rp %.0f + Rp %.0f\n", biayaKg, tambahanBiaya)
	fmt.Printf("Total biaya: Rp %.0f\n", totalBiaya)
}
