package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1

	// Menunjuk address1 untuk diubah dari Subang ke Bandung
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println()
	// Menunjuk address1 untuk diubah isinya jadi Malang
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println()
	// Membuat tabel baru = address2 untuk dengan isian Malang (melepas relasi)
	address2 = &Address{"Solo", "Jawa Tengah", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println()
	// Menunjuk address1 untuk diubah isinya jadi Maumere
	*address3 = Address{"Maumere", "Nusa Tenggara Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println()
	// Membuat tabel data baru = address3 dengan isian Makassar (melepas relasi)
	address3 = &Address{"Makassar", "Sulawesi Selatan", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}