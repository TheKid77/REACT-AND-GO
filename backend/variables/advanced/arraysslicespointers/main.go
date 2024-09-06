package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2 ,3}
	mySlice := arr1[0:1]

	fmt.Println("Array/Slice Before Func", arr1, &arr1[0], mySlice, mySlice[0])
	changeItem(arr1)
	// changeItem(mySlice)
	fmt.Println("Array/Slice After Func", arr1, &arr1[0], mySlice, mySlice[0])
}

func changeItem(arr [3]int) {
	fmt.Println("Func Array", arr)
	arr[0] = 500
}