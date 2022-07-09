package main

import "fmt"

func main() {
	type Book struct {
		name string
		id   uint
	}

	MyBook := Book{"The GO Language", 1}

	fmt.Println(MyBook)
	fmt.Println(MyBook.name)

	// updating name
	MyBook.name = "The Go Programming Language"
	fmt.Println(MyBook)

}
