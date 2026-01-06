package main

import "fmt"

func sumAll(numbers ...int16) int16 {
	var total int16 = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(12, 13, 14, 15, 16, 17, 18, 19)
	fmt.Println("Total =", total)

	slice := []int16{20, 21, 22, 23, 24, 25}
	total = sumAll(slice...)
	fmt.Println("Total2 =", total)
}