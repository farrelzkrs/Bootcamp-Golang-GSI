package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddress(address *Address) {
	address.Country = "Indonesia"
}

type Karyawan struct {
	Name, Job string
	Salary int
}

func (kyw Karyawan) CekGaji() {
	fmt.Println(kyw.Name, "dengan jabatan", kyw.Job)
	fmt.Println("Gaji:", kyw.Salary)
}

func (kyw *Karyawan) Promote(newJob string, newSalary int) {
	kyw.Job = newJob
	kyw.Salary += newSalary
	fmt.Println("Karyawan dapat promosi")
}

func (kyw *Karyawan) TaxDown(percent int) {
	cut := kyw.Salary * percent / 100
	kyw.Salary -= cut
	fmt.Println("Pajak karyawan terpotong")
}

func main() {
	address := Address{"Subang", "Jawa Barat", ""}
	ChangeAddress(&address)

	fmt.Println(address)

	worker := Karyawan {
		Name: "Lerri",
		Job: "Backend Developer",
		Salary: 5000000,
	}

	fmt.Print("Gaji awal: ")
	worker.CekGaji()
	fmt.Println()

	worker.Promote("Supervisor Backend Dev", 5000000)

	worker.CekGaji()
	fmt.Println()

	worker.TaxDown(10)

	fmt.Print("Gaji bersih: ")
	worker.CekGaji()
	fmt.Println()
}