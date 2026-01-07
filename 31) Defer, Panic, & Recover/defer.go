package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp1(val int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / val
	fmt.Println("Result =", result)
}

func main() {
	runApp1(10)
}