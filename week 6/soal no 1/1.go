package main

import "fmt"

const pi = 3.14

func volumeTabung(r float64, t float64) float64 {
	return pi * r * r * t
}

func hitungMassa(r float64, t float64, rho float64) float64 {
	return volumeTabung(r, t) * rho
}

func main() {
	var r float64
	var t1, rho1 float64
	var t2, rho2 float64

	fmt.Print("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)

	fmt.Print("Masukkan tinggi zat cair tabung kiri : ")
	fmt.Scan(&t1)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri : ")
	fmt.Scan(&rho1)

	fmt.Print("Masukkan tinggi zat cair tabung kanan : ")
	fmt.Scan(&t2)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan : ")
	fmt.Scan(&rho2)

	massaKiri := hitungMassa(r, t1, rho1)
	massaKanan := hitungMassa(r, t2, rho2)

	selisih := massaKiri - massaKanan

	if selisih < 0.0001 && selisih > -0.0001 {
		fmt.Println("BALANCE")
	} else {
		if selisih == float64(int(selisih)) {
			fmt.Printf("%.0f\n", selisih)
		} else {
			fmt.Printf("%.2f\n", selisih)
		}
	}
}
