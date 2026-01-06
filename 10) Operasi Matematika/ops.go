package main

import "fmt"

func main() {
	var (a, b int = 10, 3)
	var (
		c = a + b
		d = a * b
	)
	fmt.Println("Penjumlahan:", c)
	fmt.Println("Pengurangan:", d)

	a*=10
	fmt.Println("Nilai a dikalikan 10:", a)

	b--
	fmt.Println("Nilai b dikurangi:", b)

	var bool = true
	fmt.Println("Nilai:", !bool)
}