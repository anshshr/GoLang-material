package main

import (
	"fmt"
)
// cons - contain only fixed size and do not contain dynamic size(slices)
func main() {
	fmt.Println("we are learning arrays now")

	//initilaization
	var num[5] int;
	fmt.Println(num)


	// decleration
	var marks = [5]float32{2.34, 45.46 , 75.68,28.68}
	fmt.Println(marks);

	// if we do not declare with any values  then if will give defau;t values according to the data types as follows
	// int - 0
	// float32 - 0
	// boolean - false
	//string - ""

	var names[6] string;

	fmt.Println(names) // it will automaticvally initialize with 6 empty string ans same for other data structures as well

	//you can check here
	fmt.Printf("names is %q \n" , names)  // %q stands for the quoted string/characters it will represent the quoted characters

	// length of the array
	fmt.Println(len(names))
}