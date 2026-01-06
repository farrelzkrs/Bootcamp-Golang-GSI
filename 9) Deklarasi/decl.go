package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPJikri NoKTP = "3576648348090009"
	var marriedStatus Married = false

	fmt.Println("No KTP Jikri =", noKTPJikri)
	fmt.Println("Married Status =", marriedStatus)
}