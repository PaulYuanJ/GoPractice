package data

import "fmt"

func EP_map() {
	map1 := map[string]string{
		"Label1": "Value",
		"Label2": "Value",
	}

	map2 := map[string]int{
		"Label1": 1,
		"Label2": 3,
		"Label3": 5,
	}

	map3 := map[string]interface{}{
		"Label1": 1,
		"Label2": "??",
		"Label3": map[string]string{"childLabel": "childValue"},
	}

	// is similar with python code: if "Label3" in map3
	if value, ok := map3["Label3"]; ok {
		print(value)
	}

	fmt.Println(map1, map2, map3)

}
