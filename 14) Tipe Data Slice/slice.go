package main

import "fmt"

func main() {
	var months = []string {
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	var slice1 = months[4:7]
	fmt.Print("Slice 1: ", slice1)
	fmt.Print(", Length: ", len(slice1))
	fmt.Println(", Capacity:", cap(slice1))

	slice1[1] = "(Awas Juni)"
	fmt.Println(months)

	var slice2 = months[3:5]
	fmt.Println("Slice 2:", slice2)

	var slice3 = append(slice2, "(Bulan Baru)")
	fmt.Println("Slice 3:", slice3)
	slice3[1] = "(Desember Lagi)"
	fmt.Println("Slice now:", slice3)

	fmt.Println("Slice 2 now:", slice2)
	fmt.Println("Months now:", months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Senin"
	newSlice[1] = "Selasa"
	fmt.Println("New slice:", newSlice)
	fmt.Println("Length:", len(newSlice))
	fmt.Println("Capacity:", cap(newSlice))

	copySlice := make([]string, len(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("Copy slice:", copySlice)
}