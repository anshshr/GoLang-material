package main

import (
	"fmt"
	"strconv"

)

// converting of one data type to another data type
func main() {

	fmt.Println("learning about the data conversion")

	// int to float
	num1 := 24
	num2 := float64(num1)
	fmt.Println("the number is ", num2)
	fmt.Printf("the type  is %T \n", num2)

	// float to int
		num3 := 24.465
		num4 := int(num3)
		fmt.Println("the number is ", num4)
		fmt.Printf("the type  is %T \n ", num4)

	// string to int 
		n := 24
		str  := strconv.Itoa(n)
		fmt.Println("the number is ", str + "34")
		fmt.Printf("the type  is %T \n ", str)

	// string to int 
		s  := "24"
		n1,_  := strconv.Atoi(s)  // it return 2 values first the integer value and second an error varaible here we ignored error by _
		fmt.Println("the number is ", n1 + 20)
		fmt.Printf("the type  is %T \n ", n1)

	// string to float
		number := "3.19";
		val , _ := strconv.ParseFloat(number ,64)
		fmt.Println("the number is ", val)
		fmt.Printf("the type  is %T \n ", val)
	
	
}
