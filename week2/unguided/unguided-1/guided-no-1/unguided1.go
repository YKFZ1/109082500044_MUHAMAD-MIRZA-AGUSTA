package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var year int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&year)
	fmt.Printf("Tahun %d adalah tahun kabisat: %t\n", year, isLeapYear(year))

}
