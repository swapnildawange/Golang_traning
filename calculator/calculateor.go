package main

import (
	simpleInterst "calculator/advanced"
	"fmt"


)


func main(){
	var P float64;
	var t float64;
	var r float64
	fmt.Println("Enter Principle Amount (rupees) :")
	fmt.Scanf("%f",&P)
	fmt.Println("Time Period involved (year) :")
	fmt.Scanf("%f",&t)
	fmt.Println("Rate of Interst (%) :")
	fmt.Scanf("%f",&r)

	A:=simpleInterst.Calculate(P,r,t)
	fmt.Printf("Simple Interst for %f Rs. Principal Amount with %f percent interst rate for %f years  is %f Rs.",P,r,t,A)
}