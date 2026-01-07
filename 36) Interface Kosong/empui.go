package main

import "fmt"

func isPrime(val int) bool {
	if val <= 1 {
		return false
	}

	for j := 2; j*j <= val; j++ {
		if val%j == 0 {
			return false
		}
	}

	return true
}

func Awas(i int) interface{} {
	if isPrime(i) {
		return "Awaass!"
	} else {
		return true
	}
}

func main() {
	for i := 1; i <= 11; i++ {
		data := Awas(i)
		fmt.Printf("i: %d, hasil: %v\n", i, data)
	}
}
