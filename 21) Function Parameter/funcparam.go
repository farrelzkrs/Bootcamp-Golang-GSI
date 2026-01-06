package main

import "fmt"

func sayHello(first string, last string, age int8) {
	fmt.Println("Hello", first, last, "(", age, "tahun )")
}

func main() {
	sayHello("Farrel", "Lerri", 21)
}