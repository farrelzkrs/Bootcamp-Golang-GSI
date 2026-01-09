package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Greetings() {
	man.Name = "Mr. " + man.Name
}

type EWallet interface {
	Name() string
	Transfer(price int)
	Balance() int
}

type Ovo struct {
	User string
	Money int
}

func (ovo *Ovo) Transfer(price int) {
	ovo.Money = ovo.Money - price
	fmt.Print(ovo.User, " berhasil transfer Rp", price, "\n")
}

func (ovo *Ovo) Balance() int {
	return ovo.Money
}

func (ovo *Ovo) Name() string {
	return ovo.User
}

func Transaction(wallet EWallet) {
	fmt.Println("Nama:", wallet.Name())
	fmt.Println("Saldo awal:", wallet.Balance())
	wallet.Transfer(30000)
	fmt.Println()
	fmt.Println("Saldo akhir:", wallet.Balance())
}

func main() {
	human := Man{"Budi"}
	human.Greetings()
	fmt.Println(human.Name)
	fmt.Println()

	dompet := &Ovo {
		User: "Lerri",
		Money: 100000,
	}

	Transaction(dompet)
}