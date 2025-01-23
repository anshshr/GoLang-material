package main

import (
	"fmt"

)

// unordered collection , stores the data in form of key values pairs
func main(){
	fmt.Println("lets learn about maps");

	marks := make( map[string]int)  // initialize a map of string --> int

	// declaring the valeus (inserting the values)
	marks["rekha"] = 30;
	marks["jaya"] = 40;
	marks["suashma"] = 90;
	marks["nirma"] = 100;

	fmt.Println(marks["rekha"]); 

	//direct decleratioon and initialization
	heights := map[string]int{
		"maahi" :20,
		"krisha" :40,
		"ronak" :2,
		"shubham" :70,
	}

	fmt.Println( "height of ronak is ", heights["ronak"]) /// reading the values

	// updating the values 
	heights["ronak"] = 56

	fmt.Println( "updated height of ronak is", heights["ronak"])

	// deleting the values
	delete(heights ,"maahi") // delete the key value pair of maahi

	// checking if a key exist or not

	val , exist := heights["pankaj"]; // by default during the acces it will return two things one is the value corresponding to key and second true or false that whether key exist or not
	
	fmt.Println("exist :" , exist, " and the val is " , val);

	// iterating the values 
	
	for index , values :=  range heights{
		fmt.Println(" the key is " , index , " the values are : " , values )
	}





}