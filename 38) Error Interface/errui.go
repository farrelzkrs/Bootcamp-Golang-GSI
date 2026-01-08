package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var angka1, angka2 int
	fmt.Print("Masukkan angka1: ")
	fmt.Scanf("%d", &angka1)
	fmt.Print("Masukkan angka2 sebagai pembagi: ")
	fmt.Scanf("%d", &angka2)

	hasil, err := Pembagian(angka1, angka2)
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}