package main

import "fmt"

// defer keyword is used to execute the statements the functions in the end , if multiple defer is used then it stores everyone in the LIFO(stack order) and then executes them in the end

func add(a , b int) int {
	return a + b;
} 

func main() {
	fmt.Println("start of the program")
	defer fmt.Println("middle of the program")
	fmt.Println("end of the program")

	data := add(5 , 6)
	defer fmt.Println(" the data is " , data)

}