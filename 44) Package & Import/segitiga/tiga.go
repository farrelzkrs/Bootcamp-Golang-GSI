package segitiga

import "fmt"

func Triangle(bil int) {
	for i := 1; i <= bil; i++ {
		for j := 1; j <= bil-i; j++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}