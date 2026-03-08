package main

import (
	"fmt"
)

func main() {
	
	merah := "merah"
	kuning := "kuning"
	hijau := "hijau"
	ungu := "ungu"

	
	var p1, p2, p3, p4 string
	var p5, p6, p7, p8 string
	var p9, p10, p11, p12 string
	var p13, p14, p15, p16 string
	var p17, p18, p19, p20 string

	
	fmt.Print("Percobaan 1 : ")
	fmt.Scan(&p1, &p2, &p3, &p4)

	
	fmt.Print("Percobaan 2 : ")
	fmt.Scan(&p5, &p6, &p7, &p8)


	fmt.Print("Percobaan 3 : ")
	fmt.Scan(&p9, &p10, &p11, &p12)

	
	fmt.Print("Percobaan 4 : ")
	fmt.Scan(&p13, &p14, &p15, &p16)

	
	fmt.Print("Percobaan 5 : ")
	fmt.Scan(&p17, &p18, &p19, &p20)

	
	berhasil := true

	
	if p1 != merah || p2 != kuning || p3 != hijau || p4 != ungu {
		berhasil = false
	}

	
	if p5 != merah || p6 != kuning || p7 != hijau || p8 != ungu {
		berhasil = false
	}


	if p9 != merah || p10 != kuning || p11 != hijau || p12 != ungu {
		berhasil = false
	}

	
	if p13 != merah || p14 != kuning || p15 != hijau || p16 != ungu {
		berhasil = false
	}

	
	if p17 != merah || p18 != kuning || p19 != hijau || p20 != ungu {
		berhasil = false
	}

	
	fmt.Printf("BERHASIL : %t\n", berhasil)
}

