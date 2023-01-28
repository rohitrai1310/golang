package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var whatTosay string

	var i int

	whatTosay = "goodby harry porter"

	println(whatTosay)

	i = 5

	println("i is set to i", i)

	whatwassaid ,whatyouwant := saySomething()

	println("I know you think", whatwassaid , whatyouwant)

}

func saySomething() (string, string) {

	return ("missing something" ), ("in your life")
}
