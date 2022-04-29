package main

// In Go, an _array_ is a numbered sequence of elements of a specific length.
//Other names for arrays are collections

import "fmt"

func main() {

	// Here we create an array 'a' that will hold exactly 5 'int's.
	//The type of elements and length are both part of the array's type.
	//By default an array is zero-valued, which for 'int's means '0'.

	var a [5]int

	fmt.Println("emp:", a)

	// We can set a value at an index using the 'array[index] = value' syntax,
	// and get a value with 'array[index]'.

	a[4] = 100

	fmt.Println("set:", a)

	fmt.Println("get:", a[4])

	// The 'len' returns the length of an array.

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("dcl:", b)

	// Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.

	var twoD [2][3]int

	for i := 0; i < 2; i++ {

		for j := 0; j < 3; j++ {

			twoD[i][j] = i + j

		}

	}

	fmt.Println("2d: ", twoD)

}

// output :
//emp: [0 0 0 0 0]
//set: [0 0 0 0 100]
//get: 100
//len: 5
//dcl: [1 2 3 4 5]
//2d: [[0 1 2] [1 2 3]]

// more details :
//Each array has two properties, Type Data and Length.
//The simplest method of defining an array in Gulang is the first example inside the variable a
//An array consisting of 5 cells, which is the Length attribute and is defined by the Type Data integer int.

//If we build the array but do not set the values, fixed values ​​are placed in the cells.
// For example, if the data type is int, 0 is placed in each cell.
//Unlike languages ​​like JavaScript, the data type inside an array must be the same.
//The 'array[index] = value'  method can be used to assign a value to a specific cell in an array.
//In golang, there is a function called len that represents the length of the array, or Length.
//You can use parentheses to set an array when defining it. For example, we have how to define an array in variable b.
// we have learned how to create a one-dimensional array, but to define a multidimensional array we can use the last example, which is a two-dimensional array.
//You can use [...]string{"penn","teller"} instead of the number of houses, and in this case when compiling the number of houses
//Are counted and replaced.
//A very important point about arrays is that if we pass an array as a parameter within a function,
//that value inside the function is a copy of the array. Pointers can be used to reference the position of the array on memory.
