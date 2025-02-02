package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// we can't use two main funtion in a same project folder becuase during running of the program it creates a executable file that runs the main function that's why i had used the two main functions in the different folders

// sha256 - can only be done encoding and  decodong cannot be done means original data cannot be retrived after encoding the data

// base 64 encoding and decoding - reversible both encoding and decoding both can be done
func main() {		

	fmt.Println("we are learning about new hasinhg algorithms")

	// implementing the sha256 hashing technique

	input := "https://www.youtube.com/watch?v=OGhQhFKvMiM&t=4299s";

	data := sha256.Sum256([]byte(input));			// sah256 implememntation

	fmt.Printf("hashes data : %x\n", data);  // %x - used for looking the hashes data


	// implemnting the encoding and decoding using ghe base-62(alphanumeric - 26 uppercase + 26 lowercasee + 10 digits )

	encoded_data := base64.StdEncoding.EncodeToString([]byte(input));

	fmt.Println("encoded data is : " , encoded_data[:10])

	decoded_data , err := base64.StdEncoding.DecodeString(encoded_data[:10])

	if err != nil{
		fmt.Println("error decoding the data" , err)
		return
	}

	fmt.Println("the ddecode data is : " , string(decoded_data));
	


}