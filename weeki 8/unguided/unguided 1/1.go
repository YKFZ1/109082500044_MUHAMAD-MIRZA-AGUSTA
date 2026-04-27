package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat  Titik
	radius int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(l Lingkaran, p Titik) bool {
	return jarak(l.pusat, p) <= float64(l.radius)
}

func main() {

	var cx1, cy1, r1 int
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 int
	fmt.Scan(&cx2, &cy2, &r2)

	var x, y int
	fmt.Scan(&x, &y)

	lingkaran1 := Lingkaran{
		pusat:  Titik{x: cx1, y: cy1},
		radius: r1,
	}

	lingkaran2 := Lingkaran{
		pusat:  Titik{x: cx2, y: cy2},
		radius: r2,
	}

	titik := Titik{x: x, y: y}

	diLingkaran1 := didalam(lingkaran1, titik)
	diLingkaran2 := didalam(lingkaran2, titik)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkungan 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkungan 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkungan 2")
	} else {
		fmt.Println("Titik di luar lingkungan 1 dan 2")
	}
}
