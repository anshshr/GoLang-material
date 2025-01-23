package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("learning about the file management")

	file, err := os.Create("ansh.txt") // creation of file
	if err != nil {
		fmt.Println("failure created file")
		return
	}
	fmt.Println("success")

	// considered as good practice even if we forget to close the file if we forget to close means the resources are not free and then the probelms like memory leakage ,  inefficient memory allocation would occur
	defer file.Close()

	// writing a file
	file.WriteString("ansh shrivastav  is a developer \nand right now learning the go langauage") // completely rewrites a file removing the original content if present

	//readng a file byte by byte( advantage :- loads the file into fixed chunks and then read it so load on the memory is less )
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)

		if err == io.EOF {
			fmt.Println("ending of the file")
			break
		}

		if err == nil {
			fmt.Println("error reading the file ", err)
			return
		}

		fmt.Println(string(buffer[:n]))

	}

	// direct way to acces the data (disadavantage :- loads the complete file into the memory and the reads the file)
	data , err := os.ReadFile("ansh.txt")
	fmt.Println(string(data))

}
