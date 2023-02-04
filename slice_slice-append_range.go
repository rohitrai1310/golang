package main 

import "fmt"

func main() {

	// creating slice x
	x:= []int{1,2,4,5}

	// appending to slice x
	x = append(x,7,8,9,0 )

	// printing slice x
	fmt.Println(x)


	// print index number and item using loop 
	for i:=0 ; i<len(x) ; i++ {

		fmt.Println(i,x[i])
	}
	
	// print index number and item using range 

	for i,v := range(x) {

		fmt.Println(i,v)
	}

	// create slice y
	y:= []int{44,55,66}

	// append slice x and y
	x = append(x,y...)

	fmt.Println(x)
}