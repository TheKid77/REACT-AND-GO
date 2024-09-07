package main

import "fmt"

func main() {
	x := 10
	y := 5.2

	fmt.Println("x is", x, "y is", y, "x / y is", float64(x)/y)

	if float64(x) > y {
		println("x is larger than y")
	}
}
