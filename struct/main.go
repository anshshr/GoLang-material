package main

import "fmt"
// struct in golang - similar to class in c++ means we can define our own data structutre of our type and nothing more

// normal structs
type Person struct {
	name string
	age int
	weight int
}

type Contact struct{
	phone_no string
}

type Address struct{
	area string
	house_no int
	street string
	state string

}

// nested structs

type Employee struct{
	person Person
	add Address
	contact Contact
}
 
func main() {
	fmt.Println(" let's learn about the structs")

	// accesing or making objects
	// 1st way
	var p1  Person;
	p1.age = 20
	p1.name = "ansh"
	p1.weight = 50

	fmt.Println("the name is " ,p1.name)
	fmt.Println("the age is " ,p1.age)
	fmt.Println("the weight is " ,p1.weight)
	fmt.Println("the person is " ,p1)

	// 2nd way
	p2 := Person{
		age: 20,
		name: "risha",
		weight: 40,
	}
	
	fmt.Println("the second person is " , p2);

	// 3rd way - using the new keyword(it will return a pointer holding the address where the object is created)

	p3 := new(Person);
	p3.age = 20
	p3.weight = 59;
	p3.name = "chitu"

	fmt.Println("the person 3 is " ,p3)

	//use of nested structs
	emp := Employee{
		person: p1,
		add: Address{
			area: "maninagar" ,
			house_no: 30,
			street: "rakhial",
			state: "gujarat",
		},
		contact: Contact{
			phone_no: "123456789",
		},
	}
	fmt.Println("the employee details are", emp)



}