package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Farrel",
		"address": "Surabaya",
	}

	// Mengubah nilai
	person["title"] = "Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Menghitung jumlah elemen
	fmt.Println("Jumlah elemen:", len(person))

	// Menghapus elemen
	delete(person, "address")
	fmt.Println(person)
	fmt.Println("Jumlah elemen:", len(person))

}