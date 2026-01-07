package main

import "fmt"

type Kotak interface {
	luasKotak() int
	tipeKotak() string
}

type Adudu struct {
	Panjang int
	Lebar int
	Tipe string
}

func (adudu Adudu) luasKotak() int {
	return adudu.Lebar * adudu.Panjang
}

func (adudu Adudu) tipeKotak() string {
	if adudu.Panjang == adudu.Lebar {
		return "Persegi"
	}
	return "Persegi Panjang"
}

func hitungLuas(kotak Kotak) {
	 fmt.Println("Tipe Kotak:", kotak.tipeKotak())
	 fmt.Println("Luas Kotak:", kotak.luasKotak())
}

type HasName interface {
	getName() string
}

type Orang struct {
	Name string
}
	
func (orang Orang) getName() string {
	return orang.Name
}

func sayHello(name HasName) {
	fmt.Println("Hello", name.getName())
}

func main() {
	farrel := Orang{Name: "Farrel"}
	sayHello(farrel)
	adudu := Adudu{Panjang: 10, Lebar: 8}
	
	hitungLuas(adudu)
}