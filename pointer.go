package main

import "log"

func main() {

	var myString string

	myString = "green"

	log.Println("this is my string", myString)

	chnageusingpointer(&myString)

	log.Println("this is my new string value", myString)

}

func chnageusingpointer(s *string) {
	log.Println("this is s value", s)
	newValue := "Red"

	*s = newValue

}

// replace to a variable use &
// create a pointer type use *
//output
//"2023/01/28 22:45:47 this is my string green
//2023/01/28 22:45:47 this is s value 0x14000104220
//2023/01/28 22:45:47 this is my new string value Red
//