package main

import (
	"encoding/json"
	"fmt"
)

// lightweigh data interchange format between multiple programming langauges

// marshalling - converting into json format
// unmarshalling - converting json back into the original form

type Person struct {
	Name   string `json : "name"`
	Age    int    `json : "age"`
	Height int    `json : "height"`
}

func main() {
	fmt.Println("reprsenetataion the json")

	p1 := Person{
		Name: "ansh shrivastav",
		Age: 20,
		Height: 30,
	}

	// marhsalling the person data
	p1_data , err := json.Marshal(p1);

	if err != nil{
		fmt.Println("error mrahhalling the data" , err);
	}

	fmt.Println(string(p1_data));


	/// unmarhsalling the data

	var p2 Person;
	e := json.Unmarshal(p1_data, &p2) // always pass the address in which you wnat to store the data

	if e != nil{
		fmt.Println("error in unmarhsalling the data")
	}

	fmt.Println("data after unmarhsalling is" , p2)


}
