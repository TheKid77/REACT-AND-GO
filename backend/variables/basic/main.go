package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i1, i2 int = 1, 2
	var b bool = true
	var s string = "Hello"
	var f1 float32 = 100.101
	var f2 float64 = 200.202
	var mix1, mix2 = 22, "string"

	var (
		bi1 = 10
		bs2 = "John"
	)

	var myInt8 int8 = 127
	var myInt int = 127
	var myChar rune = 'A'

	const myConst = 100

	fmt.Printf("a is %d and b is %d\n", i1, i2)
	fmt.Println("The two vals are", i1, "and", i2)
	fmt.Println("The boolean is", b)
	fmt.Println("The string is", s)
	fmt.Println("The 32 bit float is", f1)
	fmt.Println("The 64 bit float is", f2)
	fmt.Println("mix1 is", mix1, "mix2 is", mix2)
	fmt.Println("Block bi1 is", bi1, "bs2 is", bs2)
	fmt.Println("My const is", myConst)
	fmt.Printf("myInt8: %d, Type: %s\n", myInt8, reflect.TypeOf(myInt8))
	fmt.Printf("myInt: %d, Type: %s\n", myInt, reflect.TypeOf(myInt))
	fmt.Printf("myChar: %c, Type: %s which is an alias for int32\n", myChar, reflect.TypeOf(myChar))
}
