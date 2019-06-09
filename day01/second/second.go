package second

import (
	"fmt"
)


var Name string = "Tom"

func getName() string {

	return name
}

func GetName() string {
	for i, v := range name {
		fmt.Printf("%v,%c\n", i, v)
	}

	return name
}

var name string = "Tim"
