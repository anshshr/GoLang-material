// for more ref :- https://pkg.go.dev/strings
package main

import (
	"fmt"
	"strings"
	"unicode"
)

// we will just look at different methods of string just like other programming languages so no notes !!

func main() {
	fmt.Println("different method used in the string ")

	names := "ansh,riya,shubham,dinesh,kartik";
	data := strings.Split(names , ",")
	fmt.Println(data)

	str := "anSh ShrivAstav"
	name := strings.ToUpperSpecial(unicode.CaseRanges , str)
	fmt.Println(name)

	fruit := "banana";
	f := strings.Contains(fruit , "axa")
	fmt.Println(f);

	place := "Japan";
	p := strings.Index(place , "pan")
	fmt.Println(p);

	movies := []string{"kal" , "aaj" , "pata" , "nhi"};
	m := strings.Join(movies , " --> ")
	fmt.Println(m);

	
}