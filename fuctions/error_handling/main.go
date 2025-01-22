package main

import "fmt"

func divide(a, b int) int {
	return a / b

}

// error handling we  use the error keyword
func div(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("divison by 0 is not allowed")
	}

	return a / b , nil

}

func main() {
	fmt.Println("we are learning the erorr handling in the go lang")

	fmt.Println(divide(90, 4));

	ans , erorr := div( 3, 0);
	if erorr != nil{
		fmt.Println("error occured");
	}
	fmt.Println(ans);


	res , _ := div(4 ,0) // _ is used when we  dont want to use the result coming from any variable 
	fmt.Println(res);





}
