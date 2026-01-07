package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp2(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp2(true)
}