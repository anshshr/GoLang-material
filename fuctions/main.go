package main

import "fmt"

//simple functions
func greet(){
	fmt.Println("hello buddy let's start the functions");
}

//functions with paramater
// we need to specify the types of each paramater passed also at the end of function we need to tell the return tyoe also
func add(a int , b float32) (int){
	sum := a + int(b) // we has to do the type catsing beacuse it will not add the two numbers of different type
	return sum;
}

// if all the parameters  contain same data type then write only at the end of the function
func multiply(a , b , c int)(int) {
	return a *b * c;
}

func main() {
	fmt.Println("we are here to learn the functions in golang")
	greet()
	
	sum := add(4 ,6.78)

	fmt.Println(sum)

	fmt.Println(multiply(2,3 ,10));
}