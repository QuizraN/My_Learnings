package main

import (
	sublist "exercises/01_SubList"
	etl "exercises/02_etl"
	"fmt"
)

func main() {
	fmt.Println("Executing Main Function")
	//Declarations
	l1 := []int{1, 2}
	l2 := []int{1, 2, 3}
	charIntMap := map[int][]string{1: {"A", "B"}, 2: {"C", "D"}, 3: {"E", "F"}}

	//ex-01
	fmt.Println("result is:", sublist.Sublist(l1, l2))

	//ex-02
	fmt.Println("result is:", etl.Transform(charIntMap))
}
