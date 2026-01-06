package main

import "fmt"

func main() {
	var nilai32 int32 = 100
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("nilai int32 =", nilai32)
	fmt.Println("nilai int64 =", nilai64)
	fmt.Println("nilai int8 =", nilai8)

	var nilaiInt32 int32 = 128
	var nilaiInt64 int64 = int64(nilaiInt32)
	var nilaiInt8 int8 = int8(nilaiInt32)

	fmt.Println("nilai int32 =", nilaiInt32)
	fmt.Println("nilai int64 =", nilaiInt64)
	fmt.Println("nilai int8 =", nilaiInt8)

	var name = "Jikri"
	var e byte = name[3]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}