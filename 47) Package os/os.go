package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Apa ini")
	fmt.Println(args)

	file := "aduhai.txt"

	isian := "12345678901234567890"
	err := os.WriteFile(file, []byte(isian), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("1. File telah dibuat: %s, %d\n", isian, 20)

	err = os.Truncate(file, 4)
	if err != nil {
		panic(err)
	}

	reading, _ := os.ReadFile(file)
	fmt.Printf("2. Terpotong jadi 4 byte: %s\n", string(reading))

	err = os.Truncate(file, 0)
	if err != nil {
		panic(err)
	}

	readzero, _ := os.ReadFile(file)
	fmt.Printf("3. Terpotong semua: %s, %d\n", string(readzero), len(readzero))

	os.Remove(file)
}