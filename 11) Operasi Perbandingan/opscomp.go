package main

import "fmt"

func main() {
	var (
		name1 = "Farrel"
		name2 = "Farrel"
		name3 = "Rudi"
		value1 = 10
		value2 = 100
	)

	var result bool = name1 == name2
	fmt.Println(name1, ":", name2, "=", result)

	result = name1 == name3
	fmt.Println(name1, ":", name3, "=", result)
	
	fmt.Println(value1, ">", value2, "=", value1 > value2)
	fmt.Println(value1, "<", value2, "=", value1 < value2)
	fmt.Println(value1, "==", value2, "=", value1 == value2)
	fmt.Println(value1, "!=", value2, "=", value1 != value2)
}