package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("we are learning ow to take the input in go lang")

	// scan will no take the elements afte the space

	//  var name string;
	// fmt.Scan(&name);
	// fmt.Println("hello" , name);

	//bufio.stdin will read the entire string for the input
	// we had to create a reader instance to access it and its done

	reader := bufio.NewReader(os.Stdin)
	name , _ := reader.ReadString('r')  // here we will eneter the character after which the reading would be taopeed by the reader we had created. Generally we put \n to completely read the complete string

	fmt.Println("hello" , name);
}
