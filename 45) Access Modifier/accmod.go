package main

import (
	"akses_modif/uwaw"
	"fmt"
)

func main() {
	uwaw.Terkedjoet()
	// uwaw.kaget() // error case sensitive (private function)
	fmt.Println(uwaw.Lerri)
	// fmt.Println(uwaw.farrel) // error case sensitive (private variable)
}