package main

import "fmt"

func main() {
	var final = 98
	var absensi = 80
	var finalRes bool = final >= 80
	var absensiRes bool = absensi > 80

	var lulus bool = final >= 80 && absensi > 80
	fmt.Println("Nilai final:", finalRes)
	fmt.Println("Nilai absensi:", absensiRes)
	fmt.Println("Status kelulusan:", lulus)
}