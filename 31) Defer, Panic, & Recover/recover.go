package main

import "fmt"

func endApp1() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp3(error bool) {
	defer endApp1()
	if error {
		panic("Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp3(false)
}