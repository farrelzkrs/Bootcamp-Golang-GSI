package main

import "fmt"

func main() {
	const first string = "Farrel"
	const last = "Lerri"
	const umur = uint8(21)

	fmt.Println(first, last, "(", umur, "tahun )")

	const (
		awal = "Farrel"
		akhir = "Zikri"
		age = int8(28)
	)

	fmt.Println(awal, akhir, "(", age, "tahun )")
}