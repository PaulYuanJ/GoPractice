package data

import (
	"fmt"
	"reflect"
)

func EP_interface() {

	var x map[string]map[string]string = map[string]map[string]string{
		"1": {
			"name": "paul",
		},
		"2": {
			"name": "paul",
		},
		"3": {
			"name": "paul",
		},
	}

	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x).Field(0))
	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(reflect.ValueOf(x).MapKeys())
	fmt.Println(reflect.ValueOf(x).MapRange())

}
