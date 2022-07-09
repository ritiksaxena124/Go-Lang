package main

import "fmt"

func main() {
	people := map[string]int{}
	people["Raju"] = 6
	fmt.Println(people)

	students := map[int]string{
		1: "Raju",
		2: "Kaju",
		3: "Laju",
	}
	fmt.Println(students)

	// print mapping using for loop
	for key, value := range students {
		fmt.Println(key, value)
	}
}
