package main

import "fmt"

func main() {
	name := "Farrel"

	// Normal Switch
	switch name {
	case "Farrel":
		fmt.Println("Hello Farrel")
	case "Doni":
		fmt.Println("Hello Doni")
	default:
		fmt.Println("Hello User")
	}

	// Switch with Expression
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

	// Switch Without Expression
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama masih panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama udah benar")
	}
}