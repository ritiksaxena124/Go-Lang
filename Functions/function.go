package main

import "fmt"

func sum(a int, b int) int {
	res := a + b
	return res
}

func main() {
	var ans int = sum(10, 20)
	fmt.Println(ans)
}
