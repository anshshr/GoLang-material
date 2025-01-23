package main

import (
	"fmt"

)

func main() {
	fmt.Println("now switch case");

	// normal switch case
	age := 45;

	switch age{
	case 18:
		fmt.Println("age is 18")
	case 20:
		fmt.Println("age is 20")
	case 40:
		fmt.Println("age is 40")
	case 60:
		fmt.Println("age is 60")
	case 80:
		fmt.Println("age is 80")
	default:
		fmt.Println("nahi pata bhai")
	}

	// can use multilple values as well
	switch age{
	case 1, 19 ,46, 36:
		fmt.Println("age is betwen 1 and 46")
	case 20, 67 ,78:
		fmt.Println("age is between 20 and 78")
	case 40:
		fmt.Println("age is 40")
	case 60:
		fmt.Println("age is 60")
	case 80:
		fmt.Println("age is 80")
	default:
		fmt.Println("nahi pata bhai")
	}



	// can use the epression as well ans also we dont need to pass the deciding parameter in switch
	switch {
	case  age >= 0 &&  age < 18:
		fmt.Println("not eligible for vote");
	case age >= 18 && age < 60:
		fmt.Println("eligible for vote");
	default:
		fmt.Println("adult voters");

	
	}


}