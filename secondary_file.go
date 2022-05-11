package main

//In Go, an array is a numbered sequence of elements of a specific length.
//Here we create an array a that will hold exactly 5 ints.
//The type of elements and length are both part of the array’s type.
// By default an array is zero-valued, which for ints means 0s.
//We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
//The builtin len returns the length of an array.
//Use this syntax to declare and initialize an array in one line.
//Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
//Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.

import (
	"fmt"
)

func main() {

	myArray := [...]string{"Apples", "Oranges", "Bananas"}

	fmt.Printf("Initial array values: %v\n", myArray)

	myFunction(myArray)

	fmt.Printf("Final array values: %v\n", myArray)

}

func myFunction(arr [3]string) {

	// Change Oranges to Strawberries

	arr[1] = "Strawberries"

	fmt.Printf("Array values in myFunction(): %v\n", arr)

}

// Output
// Initial array values: [Apples Oranges Bananas]
// Array values in myFunction(): [Apples Strawberries Bananas]
// Final array values: [Apples Oranges Bananas]

//By printing an array, a value like [v1,v2,v3,...] is displayed.
//The important thing about arrays is that we can not set the number of cells more than the specified value,
//meaning the array is not flexible.
