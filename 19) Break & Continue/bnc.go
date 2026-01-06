package main

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i == 7 {
			break
		}

		println("Perulangan ke", i)
	}
}