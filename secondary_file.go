package main

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
