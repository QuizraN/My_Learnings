package main

import (
	sublist "exercises/01_subList"
	etl "exercises/02_etl"
	school "exercises/03_grades"
	interfaces "exercises/04_interfaces"
	jsondata "exercises/05_jsonData"
	"fmt"
)

func main() {
	fmt.Println("Executing Main Function")
	//Declarations
	l1 := []int{1, 2, 1, 2, 3}
	l2 := []int{1, 2, 3}
	charIntMap := map[int][]string{1: {"A", "B"}, 2: {"C", "D"}, 3: {"E", "F"}}

	//ex-01
	fmt.Println("ex-01 result is:", sublist.Sublist(l1, l2))

	//ex-02
	fmt.Println("ex-02 result is:", etl.Transform(charIntMap))

	//ex-03
	fmt.Println("ex-03 result is:")
	school.TestSchool()

	//ex-04
	fmt.Println("ex-04 result is:")
	interfaces.Display()

	//ex-05
	fmt.Println("ex-05 result is:")
	// jsondata.GetAllUser()
	jsondata.CreateUser()

}
