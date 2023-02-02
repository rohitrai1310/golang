package main 

import "fmt"

func main() {

	// 2-D int array s
	s:= [2][2]int{ {2 ,3} , {4,5} }

	//var s [2][2]int
	//s[0][0] = 100
    //s[0][1] = 200
    //s[1][0] = 300
    //s[1][1] = 400


	//Printing 2-D int array 

	for i:=0 ; i<2 ; i++ {
		for j:=0 ; j<2 ; j++ {
			fmt.Println(s[i][j])

		}
	}
}