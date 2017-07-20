package main

import "fmt"

func main() {


	//https://blog.golang.org/go-slices-usage-and-internals


	//TOPICS:
	//slices can be based on an array or directly from a slice
	//a slice can be declared with the make function
	//when the pointer moves you loose any data that was before the pointer
	//to increase size of slice have to copy data
	//can also use the append
	//appending overlapping slices
	//appending two slices together


	a := [7]int{3,1,7,8,4}
	fmt.Println("an array length of 7 with 5 values")
	fmt.Println(a)

	//declare new slice
	s := []int{}

	fmt.Println("a slice created based on an array")
	s = a[0:7]
	fmt.Println(s)

	////commented out because causes compile error
	//fmt.Println("try to make slice bigger")
	//s = a[0:12]
	//fmt.Println("can't make slice larger than array it is based on without using make()")

	fmt.Println("use make to create a larger slice")
	s = make([]int, 0, 12)
	printSlice(s)

	s = []int{2, 3, 5, 7, 11, 13}
	fmt.Println("initialize")
	printSlice(s)

	// Slice the slice to give it zero length.
	//fmt.Println("slice to zero length")
	//s = s[:0]
	//printSlice(s)

	// Extend its length.
	fmt.Println("extend its length")
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	fmt.Println("Drop its first two values")
	s = s[2:]
	printSlice(s)

	//have all values back again
	fmt.Println("add all back")
	s = s[0:4]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d addr=%p %v\n", len(s), cap(s), &s, s)
}