package main

import "fmt"

func main() {
	// ways to declare multiple var
	name, age := "Ritik", 21
	fmt.Println(name, age)
	// updating name value
	name, age1 := "Ritik Saxena", 21
	fmt.Println(name, age1)

	var (
		school  string
		roll_no int
	)

	fmt.Println(school, roll_no)
}
