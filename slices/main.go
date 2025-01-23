package main

import "fmt"

// slices have the dynammic size(internal implementation is array)
func main() {
	fmt.Println("we are now learning the slices - dynamic size ")

	// decleartion of an empty slice
	numbers := []int{}
	fmt.Println(numbers)

	numbers = append(numbers, 2, 4, 5, 6, 4) // similar to push_back of vector

	fmt.Println(numbers)
	fmt.Println(len(numbers)) // for length of the slice
	fmt.Println(cap(numbers)) // for caapcity of the slice

	// make function - used to define the slice with a specific length and caapcity

	//  -> data type , length of slice , capacity of the slice
	stock := make([]int, 0, 0) // empty slice of length and capacity 0
	fmt.Println(stock)

	// capcacity - if we insert the elements in a slice greater than the capacaity the the capacity automatically doubles the original value to store the values that's why it is dynamic in size
	marks := make([]int, 4, 5)
	marks = append(marks, 2, 3, 4, 5)

	fmt.Println("slice :", marks)
	fmt.Println("len :", len(marks))
	fmt.Println("capacity :", cap(marks))
}
