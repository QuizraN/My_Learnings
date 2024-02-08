package main

import (
	sublist "exercises/01_SubList"
	"fmt"
)

func main() {
	fmt.Println("Executing Main Function")
	//ex-01
	l1 := []int{1, 2}
	l2 := []int{1, 2, 3}
	result := sublist.Sublist(l1, l2)
	fmt.Println("result is:", result)
}
