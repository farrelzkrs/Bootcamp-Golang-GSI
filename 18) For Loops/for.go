package main

import "fmt"

func main() {
	for cnt2 := 1; cnt2 <= 5; cnt2++ {
		fmt.Println("Perulangan ke", cnt2)
	}
	fmt.Println()

	names := []string{"Farrel", "Doni", "Rizki", "Ahmad", "Budi"}
	for i, name := range names {
		fmt.Println("Nomor", i+1, "=", name)
	}
}