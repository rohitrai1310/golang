package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%d is room number of %s and it is %t", x, y, z)
	s1 := fmt.Sprintf("%v\t , %v\t , %v\t", x, y, z)
	fmt.Println(s)
	fmt.Println(s1)
}

//42 is room number of James Bond and it is true
//42       , James Bond    , true
