package main

import "fmt"

func getFullName() (string, string, string) {
	return "Farrel", "Lerri", "Putra"
}

func main() {
	first, mid, last := getFullName()
	fmt.Println(first, mid, last)

	_, two, _ := getFullName()
	fmt.Println(two)
}