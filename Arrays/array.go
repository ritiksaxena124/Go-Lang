package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	arr[0] = 10        // change value at particular index
	fmt.Println(arr)   // print array
	length := len(arr) // returns length of array
	// print array using loop
	for i := 0; i < length; i++ {
		fmt.Println(arr[i])
	}

}
