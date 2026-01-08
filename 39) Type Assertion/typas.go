package main

import "fmt"

type JaringanError struct {
	Pesan string
}

type ValidError struct {
	Field string
}

func (e *JaringanError) Error() string {
	return e.Pesan
}

func (e *ValidError) Error() string {
	return "Salah input di: " + e.Field
}

func simsError(tipe string) error {
	if tipe == "dc" {
		return &JaringanError{
			Pesan: "You are disconnected...",
		}
	}
	return &ValidError{
		Field: "Email atau password salah",
	}
}

func random() interface{} {
	return false
}

func main() {
	// var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// Switch case mendeteksi jenis nilai return
	hasil := random()
	switch nilai := hasil.(type) {
	case string:
		fmt.Println("String", nilai)
	case int:
		fmt.Println("Integer", nilai)
	default:
		fmt.Println("Yntkts")
	}

	// Simulasi Error
	err := simsError("dc")
	switch e := err.(type) {
	case *JaringanError:
		fmt.Println("Ngelag bang:", e.Pesan)
	case *ValidError:
		fmt.Println("Cek lagi masukannya:", e.Field)
	}
}