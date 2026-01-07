package main

import "fmt"

type Manusia struct {
	Name, Address string
	Age int8
	Married bool // Secara default bernilai false
}

func main() {
	var man Manusia
	man.Name = "Farrel Zikri"
	man.Age = 20
	man.Address = "Surabaya"
	fmt.Println(man)

	lerri := Manusia{
		Name: "Lerri",
		Age: 32,
		Address: "Bandung",
	}
	fmt.Println(lerri)

	// jikri := Manusia{"Jikri", "Malang", 28}
}