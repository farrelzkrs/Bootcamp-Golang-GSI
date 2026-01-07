package main

import "fmt"

func main() {
	name := "Farrel"
	count := 0
	incerement := func() {
		name := "Budi"
		fmt.Println("Incerement")
		fmt.Println(name)
		count++
	}

	incerement()
	incerement()

	fmt.Println(count)
	fmt.Println(name)
}