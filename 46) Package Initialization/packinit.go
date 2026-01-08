package main

import (
	"cobadata/database"
	"cobadata/sinus"
	"fmt" /* _ (blank indentification) */
)

func main() {
	result := database.GetData()
	fmt.Println(result)

	fmt.Println("Ini Tabel Sinus")
	sudut := []int {0, 15, 30, 45, 60, 75, 90}

	for _, sdt := range sudut {
		nilai := sinus.TakeTable(sdt)
		fmt.Printf("Sin %d adalah %.2f\n", sdt, nilai)
	}
}