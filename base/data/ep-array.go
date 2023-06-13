package data

import "fmt"

func EP_array() {

	// create
	// Output: [One Two Three]
	// OutType: array
	myStringArray := []string{
		"One",
		"Two",
		"Three",
	}
	fmt.Println(myStringArray)

	// retrieve the second element in the array
	// Output: Two
	// OutType: string
	fmt.Println(myStringArray[1])

	// retrieve the second to third element in the array
	// Output: [Two]
	// OutType: slice
	fmt.Println(myStringArray[1:2])

	// modify
	// Output: [One Modified_One Three]
	// OutType: array
	myStringArray[1] = "Modified_One"
	fmt.Println(myStringArray)

	// delete
	// Output: [One Modified_One Three]
	// OutType: array
	myStringArray[1] = "Modified_One"
	fmt.Println(myStringArray)



	// TODO more examples will come soon...

}