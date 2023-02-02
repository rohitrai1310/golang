package main

import "fmt"

func main () {

	// define staring s
	s:= [...]int{1,2,3,4,5,6,6,7}

	// Print length of array s

	fmt.Println("length of array is: " ,len(s))
	/// Print integer 
	fmt.Println("this is the value fo integer")

	for i:=0 ; i<len(s) ; i++ {

		fmt.Println(s[i])
	} 
}