package main

import "fmt"

func factor(val int) int {
	result := 1
	for i := val; i > 0; i-- {
		result *= i
	}
	return result
}

func recur(val int) int {
	if val == 1 {
		return 1
	} else {
		return val * recur(val - 1)
	}
}

func fibo(val int) int {
	if val <= 1 {
		return val
	} else {
		return fibo(val - 1) + fibo(val - 2)
	}
}

func main() {
	val := 5
	fibfib := 12
	fmt.Println("Loop =", factor(val))
	fmt.Println("Recursive =", recur(val))
	fmt.Println("Manual =", 5 * 4 * 3 * 2 * 1)
	fmt.Println()
	fmt.Println("Fibonacci =", fibo(fibfib))

}