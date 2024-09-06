package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2 ,3}
	var arr2 = [...]int{1, 2 ,3}
	var empty = [...]int{}
	arr3:= [...]int{1, 2 ,3}
	fmt.Println("Arrayy is", arr1, arr2, arr3)
	for i:=0;i < len(arr1);i++ {
		fmt.Println("Arr Element", i, "is", arr1[i])
	}
	fmt.Println("Empty Array", empty, len(empty), cap(empty) )
}