package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	f := func() int {
		s := a
		a = b
		b = s + b
		return s
	}
	return f
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
