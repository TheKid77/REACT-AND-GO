package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2 ,3}
	mySlice := arr1[0:1]
	fmt.Println("The slice length is" , len(mySlice))
	fmt.Println("The slice capacity is" , cap(mySlice))
	for i:=0;i < len(mySlice);i++ {
		fmt.Println("Slice", i, "is", mySlice[i])
	}
}