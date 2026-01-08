package main

import (
	"cossinus/cosinus"
	"fmt"
)

func main() {
	fmt.Println("Ini Tabel Cosinus")
	sudut := []int {0, 15, 30, 45, 60, 75, 90}

	for _, sdt := range sudut {
		nilai := cosinus.TakeTable(sdt)
		fmt.Printf("Cosinus %d adalah %.2f\n", sdt, nilai)
	}
}