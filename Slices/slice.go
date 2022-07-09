package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	fmt.Println(slice)

	slice[0] = 21 // change element at particular index

	fmt.Println(slice)

	slice1 := append(slice, 100) // append element at the end of our slice
	fmt.Println(slice, slice1)

}
