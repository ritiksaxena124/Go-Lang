package main

import "fmt"

func main() {
	// Strings are immutable in GO
	var str string = "Hello World!"
	fmt.Println(str)
	// concatenating two strings
	fmt.Println(str + "This is my first GO Program")
	fmt.Println(str[0])   // 	Prints ASCII value of char at that index
	fmt.Println(str[2:7]) // slicing: llo W
}
