package main

import "fmt"

// pointers -> generally use to store the address of another varaible and nothing more easy peasy
// used for efficient memeory management and data management between functions

func main() {
	fmt.Println("looking into the pointers ")

	num := 2;
	var ptr *int = &num ; // pointers variables should be of same data type as of which varaibale address they are going to contain
	
	//another way
	ptr2 := &num // this will automatically find that which data address we are giving it to store and automatically make a pointer for that

	fmt.Println("the address is " , ptr)
	fmt.Println("the address is " , ptr2);
	fmt.Println("the data is" ,*ptr)


	var pointer * int;
	fmt.Println(pointer) // it will hold no data
}