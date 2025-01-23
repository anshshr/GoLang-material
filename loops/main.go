package main

import (
	"fmt"

)

// only contain for loop and not dowhile or while loop

func main() {
	fmt.Println("lets study about the loops")
	

	// normal loop( V.V.IMP --> do not use the braces in the loop else it will throw the  error)

	for i := 0 ; i < 4 ; i++{
		fmt.Println("the numbers is " , i);
	}

	// iteration on slices and string

	name := "ansh shrivasatav"
	for i := 0 ; i < len(name) ; i++{
		fmt.Println(name[i] ) // this will return the ascii values
		fmt.Printf("the values is %c \n" , name[i]) // this will return the actual character
	}

	number := []int{4,3,2,3,4,4,34,3,4};
	for i := 0 ; i < len(number) ; i++{
		fmt.Println(number[i])
	}

	// range in loops - will give the index as well as the values
	for index , values := range name{
		fmt.Println("the index is " ,index , " the characters is " , values) // it will still  give the ascii values

		fmt.Printf("the values is %c  with the index as %d \n" , values , index) // this will give the comlpete character
	}

}