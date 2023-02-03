package main

import "fmt"

func main() {

	// 1-D array example with short declaration variable

	x := [4]string{"A", "B", "C", "D"}

	// Printing value of x
	fmt.Println("Printing value of x\n")

	for i := 0; i <= 3; i++ {

		//fmt.Println(x[i])
		fmt.Printf("%#s", x[i])
	}

	// 1-D creating array using var keyword
	var y [3]int

	y[0] = 1
	y[1] = 2
	y[2] = 3

	// Printing value of y
	fmt.Println("\n\nPrinting value of y")

	for i := 0; i <= 2; i++ {

		fmt.Println(y[i])

	}

	// multi dimension short declaration array example

	array1 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println("\nPrint array1")

	for i := 0; i <= 2; i++ {

		for j := 0; j <= 2; j++ {

			fmt.Println(array1[i][j])
		}
	}

	// multi dimension var dec;aration array example

	var array2 [3][3]string

	array2[0][0] = "A"
	array2[0][1] = "B"
	array2[0][2] = "C"
	array2[1][0] = "D"
	array2[1][1] = "E"
	array2[1][2] = "F"
	array2[2][0] = "G"
	array2[2][1] = "H"
	array2[2][2] = "I"

	fmt.Println("\n\nPrint value of array2")

	for i := 0; i <= 2; i++ {

		for j := 0; j <= 2; j++ {

			fmt.Println(array2[i][j])
		}
	}

}

/* OUTPUT

Printing value of x

ABCD

Printing value of y
1
2
3

Print array1
1
2
3
4
5
6
7
8
9


Print value of array2
A
B
C
D
E
F
G
H
I
*/
