package main

import "fmt"

func main() {
	fmt.Println("hello world");

	name := "ansh srivastav";
	age := 19
	weight := 50.652387792

	//for customised input we use the print f
	fmt.Printf("name : %s , age : %d , weight : %.4f \n" , name , age , weight);

	// println atomativally  adds thr spaces between the elemnts so we would no be able to customise it
	fmt.Println("name is :", name ,"age is :" ,age ,"height is :" , weight)
}