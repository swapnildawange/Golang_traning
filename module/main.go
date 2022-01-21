package main

import (
	"fmt"

	sliceOperations "github.com/sliceOperations/sliceOperations" // importing from local 
	// sliceOperations "github.com/swapnildawange/sliceOperations" // importing from github
)

func execute() {
	s1 := []string{"a", "b", "c", "d"}
	s2 := []string{"a", "b", "g", "k"}

	ele := "b"
	isPresent := sliceOperations.Contains(ele, s1)

	if isPresent {
		fmt.Printf("%v is Present in slice %v\n", ele, s1)
	} else {
		fmt.Printf("%v is Not Present in slice %v\n", ele, s1)
	}

	diff := sliceOperations.Difference(s1, s2)

	if len(diff) > 0 {
		fmt.Printf("%v are present in slice %v but not in %v \n", diff, s1, s2)
	} else {
		fmt.Println("No elements found")
	}

	unique := sliceOperations.Unique(s1)
	if len(unique) > 0 {
		fmt.Printf("Unique element in %v are %v \n", s1, unique)
	} else {
		fmt.Println("No unique elements found")
	}
}

func main() {
	execute()
}
