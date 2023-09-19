package data

import (
	"fmt"
	"time"
)
import "github.com/patrickmn/go-cache"

var goCahce *cache.Cache //定义全局变量

func init() {
	goCahce = cache.New(5*time.Minute, 60*time.Second)
}

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
	myStringArray = append(myStringArray[:1], myStringArray[2:]...)
	fmt.Println(myStringArray)

	// TODO more examples will come soon...

	var temp []string

	myCommonArray := []string{"a", "b", "c", "d", "e"}

	for _, i := range myCommonArray {
		temp = append(myCommonArray[:0], myCommonArray[1:]...)
		fmt.Println(i, temp)
	}

}
