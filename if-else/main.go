package main

import (
	"fmt"

)

func main() {
	fmt.Println("now its turn for the conditions statements")

	//simple and sweet 

	// you can put braces in condition or not its totally up to you   go will automatically whether there is a need for braces or not and remove that if not necessary
	name := "camalcase"

	if(name == "camalcase"){
		fmt.Println(name)
	}else{
		fmt.Println("no camalcase")
	}

	// if else ladder
	num := 7;

	if(num >5 ){
		fmt.Println("number is greater than 5")
	}else if(num < 10){
		fmt.Println("number is less than 10")
	}else if(num <5){
		fmt.Println("number is less than 5")
	}else{
		fmt.Println("i ma the else condition")
	}
}
