package main

import (
	"fmt"
	"strings"
)

func main() {
	clone := strings.Clone("Budi")
	contains := strings.Contains("Farrel Lerri", "Rel")
	lower := strings.ToLower("KAGET")
	upper := strings.ToUpper("lerri")
	replace := strings.Replace("Farrel Zikri Lerri Lerri", "Lerri", "Suryahadi", 1)
	fmt.Println(clone)
	fmt.Println(contains)
	fmt.Println(lower)
	fmt.Println(upper)
	fmt.Println(replace)


}