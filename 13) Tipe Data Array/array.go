package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Farrel"
	names[1] = "Zikri"
	names[2] = "Lerri"

	fmt.Println("Nama pertama:", names[0])
	fmt.Println("Nama kedua:", names[1])
	fmt.Println("Nama ketiga:", names[2])

	var values = [3]int{90, 80, 70}
	fmt.Println("Nilai:", values)
	values[0] = 100
	values[1] = 70
	fmt.Println("Nilai setelah diubah:", values)

	fmt.Println(len(names))
	fmt.Println(len(values))
}