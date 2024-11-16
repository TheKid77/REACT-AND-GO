package main

import (
	"fmt"
)

func subslices(x []int) {
	fmt.Println("SLICING SLICES")
	fmt.Println(x)
	x[0] = 1
	subslice1 := x[:3]
	subslice1[1] = 10
	fmt.Println(subslice1)
	var xy = []int{1, 2, 3}
	subslice := xy[:2]
	subslice[1] = 5
	fmt.Println(subslice)
	clear(subslice)
	fmt.Println(subslice)
	clear(x)

}

func makeslices() {
	fmt.Println("USING MAKES")
	x := make([]int, 10, 10)
	fmt.Printf("make[]int, 10, 10 produces %v, len of %d, capacity of %d\n", x, len(x), cap(x))
	x = append(x, 1, 2)
	fmt.Printf("after append(x, 1) gives %v, len of %d, capacity of %d\n", x, len(x), cap(x))
	y := x[:]
	y[0] = 1
	fmt.Println(x, y)
	z := make([]int, 10)
	copy(z, y)
	z[0] = 2
	fmt.Println(x, y, z)
}

func demo() {
	fmt.Println("DEMO")
	a := []int{1, 2, 3}
	y := a[:]
	a[0] = 4
	fmt.Println(y)
	z := make([]int, 10)
	a[0] = 1
	copy(z, y)
	y[0] = 4
	fmt.Println(z)
}
func main() {
	var slice1 []int
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 10)
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 20)
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 30)
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 40)
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 50)
	fmt.Println(slice1, len(slice1), cap(slice1))
	var slice2 = []int{1, 2, 5: 3}
	fmt.Println(slice2, len(slice2), cap(slice2))
	var slice3 = make([]int, 5, 10)
	fmt.Println(slice3, len(slice3), cap(slice3))
	slice3 = append(slice1, 1, 2, 3, 4)
	fmt.Println(slice3, len(slice3), cap(slice3))
	clear(slice3)
	fmt.Println(slice3, len(slice3), cap(slice3))

	subslices(slice3)
	fmt.Println("BACK IN MAIN")
	fmt.Println(slice3, len(slice3), cap(slice3))
	makeslices()
	demo()
	fmt.Printf("The slice capacity is %s")
}
