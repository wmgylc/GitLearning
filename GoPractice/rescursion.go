package main

import "fmt"

func FibonacciSequence(a, b int) {
	if a == b {
		fmt.Println(1)
	}
	fmt.Println(b)
	var tmp = b
	b += a
	a = tmp
	if b < 1000 {
		FibonacciSequence(a, b)
	}

}

func main() {
	var a, b = 1, 1
	FibonacciSequence(a, b)
}
