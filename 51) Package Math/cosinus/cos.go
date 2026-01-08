package cosinus

import (
	"fmt"
	"math"
)

var table map[int] float64

func init() {
	fmt.Println("Tabel Cosinus")
	table = make(map[int] float64)

	for i := 0; i < 360; i+= 15 {
		radius := float64(i) * (math.Pi / 180)
		table[i] = math.Cos(radius)
	}
}

func TakeTable(sudut int) float64 {
	return table[sudut]
}