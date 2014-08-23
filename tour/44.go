package main

import "fmt"

func fibonacchi() func() int {
	f, g := 0, 1
	return func() int {
		f, g = g, f+g
		return f
	}
}

func main() {
	f := fibonacchi()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
