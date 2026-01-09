package main

import (
	"fmt"
	"strconv"
)

func main() {
	atoi, _ := strconv.Atoi("1234567890")
	itoa := strconv.Itoa(1234567890)

	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	fmt.Println(atoi)
	fmt.Println(itoa)

}
