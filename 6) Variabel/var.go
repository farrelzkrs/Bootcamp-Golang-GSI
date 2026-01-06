package main

import (
	"fmt"
)

func main() {
	var name1 string

	name1 = "Farrel Zikri"
	fmt.Println(name1)

	name1 = "Farrel Lerri"
	fmt.Println(name1)

	var name2 = "Lerri"
	fmt.Println(name2)

	var umur = uint8(21)
	fmt.Println(umur)

	rumah := "Surabaya"
	fmt.Println(rumah)

	var (
		firstName = "Farrel"
		lastName = "Lerri"
	)

	fmt.Println(firstName, lastName)
}