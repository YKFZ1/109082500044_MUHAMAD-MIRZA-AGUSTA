package main

import "fmt"

type suhu float64

func CelciusToReamur(Celcius suhu) suhu {
	return suhu(float64(Celcius) * 4 / 5)
}

func CelciusToFahrenheit(Celcius suhu) suhu {
	return suhu(float64(Celcius)*9/5 + 32)
}

func CelciusToKelvin(Celcius suhu) suhu {
	return suhu(float64(Celcius) + 273.15)
}

func main() {
	var celcius suhu

	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scan(&celcius)

	fmt.Println("\n=== HASIL KONVERSI ===")
	fmt.Printf("Celsius: %.2f°C\n", celcius)
	fmt.Printf("Reamur: %.2f°R\n", CelciusToReamur(celcius))
	fmt.Printf("Fahrenheit: %.2f°F\n", CelciusToFahrenheit(celcius))
	fmt.Printf("Kelvin: %.2f K\n", CelciusToKelvin(celcius))
}
