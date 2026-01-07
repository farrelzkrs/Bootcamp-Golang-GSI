package main

import (
	"fmt"
	"time"
)

type Manungsa struct {
	Name string
	Job string
	Age int8
}

func (manungsa Manungsa) helloBang(name string) {
	fmt.Println("Hello", name + ", my name is", manungsa.Name)
}

func main() {
	var ler Manungsa
	cnt := 1
	for i := 3; i >= cnt; i-- {
		fmt.Println(i)
		time.Sleep(time.Duration(1 * time.Second))
	}
	ler.Name = "Farrel"
	ler.helloBang("Lerri")
}