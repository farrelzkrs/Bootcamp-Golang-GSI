package main

import "fmt"

type Filter func(string, string) string

func helloFilter(name string, last string, filter Filter) {
	nameFilter := filter(name, last)
	fmt.Println("Hello ", nameFilter)
}

func kataKasarFilter(name string, last string) string {
	if name == "Kasar" || last == "Kasar" {
		if name == "Kasar" {
			return "..." + " " + last
		} else {
			return name + " " + "..."
		}
	} else {
		return name + " " + last
	}
}

func main() {
	helloFilter("Kasar", "Lerri", kataKasarFilter)
	helloFilter("Farrel", "Kasar", kataKasarFilter)
}