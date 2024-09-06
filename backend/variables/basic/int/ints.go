package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt8 int8 = 127
	var myInt int = 127
	fmt.Printf("myInt8: %d, Type: %s\n", myInt8, reflect.TypeOf(myInt8))
	fmt.Printf("myInt: %d, Type: %s\n", myInt, reflect.TypeOf(myInt))

}
