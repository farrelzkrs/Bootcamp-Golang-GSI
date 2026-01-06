package main

import "fmt"

func main() {
	name := "Doni"

	if name == "Farrel" {
		fmt.Println("Hello Farrel")
	} else if name == "Doni" {
		fmt.Println("Hello Doni")
	} else {
		fmt.Println("Hello User")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}