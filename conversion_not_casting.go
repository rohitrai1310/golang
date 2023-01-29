package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 20
	fmt.Println("value of a is", a)
	fmt.Printf("%T\n", a)

	b = 55

	fmt.Println("this is the value of b", b)

	fmt.Printf("%T\n", b)

	//a = b
	//fmt.Println("this is new a", a)

	//fmt.Printf("%T\n", a)
	//cannot use b (variable of type hotdog) as type int in assignment

	a = int(b)

	fmt.Println("this is the new value of a", a)
	fmt.Printf("%T\n", a)

}
