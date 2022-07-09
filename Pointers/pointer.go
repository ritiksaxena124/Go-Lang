package main

import "fmt"

func main() {
	a := 2
	fmt.Println(&a) // Prints address of a in the memory

	// ptr is a Pointer
	ptr := &a // ptr stores address of a
	fmt.Println(ptr)

	// Dereferencing
	fmt.Println(*ptr) // Prints value stored at the address that the ptr pointer is storing

	fmt.Println(&ptr) // Prints address of pointer ptr
}
