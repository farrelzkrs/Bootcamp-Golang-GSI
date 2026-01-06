package main

import "fmt"

func getName() (first string, middle string, last string) {
	first = "Farrel"
	middle = "Lerri"
	last = "Putra"
	return
}

func main() {
	f, m, l := getName()
	fmt.Println(f, m, l)

	_, mid, last := getName()
	fmt.Println(mid, last)
}