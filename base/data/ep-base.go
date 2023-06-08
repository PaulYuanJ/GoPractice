/*
This package provides golang base practice examples.
string
int
float
bool
array
slice
map
struct
channel
*/
package data

import "fmt"

func EP_string() {
	// define string variables

	// use := to define
	myString1 := "Hello World!!"
	fmt.Println(myString1)

	// define the variable type first
	var myString2 string
	myString2 = "Hello Golang!"
	fmt.Println(myString2)

	// use fmt.Sprintf to contact to strings
	fmt.Println(fmt.Sprintf("%s - %s", myString1, myString2))

	// TODO more examples will come soon...

}
