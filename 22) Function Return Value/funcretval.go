package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bang"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Lerri")
	fmt.Println(result)
	fmt.Println(getHello(""))
}