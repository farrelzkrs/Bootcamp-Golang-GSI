package main

import "fmt"

type Blacklist func(string) bool

func registerUser(orang []string, blacklist Blacklist) {
	for i := 0; i < len(orang); i++ {
		if blacklist(orang[i]) {
			fmt.Println("Kamu di block,", orang[i])
		} else {
			fmt.Println("Welcome back,", orang[i])
		}
	}
}

func main() {
	manusia := []string{"admin", "Lerri", "Budi", "root", "anonymous"}
	registerUser(manusia, func(orang string) bool {
		if orang == "admin" || orang == "root" || orang == "anonymous" {
			return true
		} else {
			return false
		}
	})
}