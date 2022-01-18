package main

// Write a program to calculate the sum of squares of numbers 1 to 1000, using 5 goroutines.

import (
	"fmt"
	"sync"
)

type digits struct {
	number int
	arr    []int
}

type squareSum struct {
	number int
	sum    int
}

func getDigits(number int, digitsStruct chan digits, wg *sync.WaitGroup) {
	var digitsArr []int
	originalNumber := number
	for number != 0 {
		currentDigit := number % 10
		number = number / 10
		digitsArr = append(digitsArr, currentDigit)
	}
	digitsStruct <- digits{originalNumber, digitsArr}

	wg.Done()

}

func getSquare(digitsChan chan digits, squaresChan chan digits, wg *sync.WaitGroup) {
	var squaresArr []int
	digitsStruct := <-digitsChan
	for _, i := range digitsStruct.arr {
		numSquare := i * i
		squaresArr = append(squaresArr, numSquare)

	}
	squaresChan <- digits{number: digitsStruct.number, arr: squaresArr}
	wg.Done()
}

func getSum(squaresChan chan digits, sumChan chan squareSum, wg *sync.WaitGroup) {
	squaresStruct := <-squaresChan
	sum := 0
	for _, i := range squaresStruct.arr {
		sum = sum + i
	}

	sumChan <- squareSum{number: squaresStruct.number, sum: sum}

	wg.Done()
}
func printSum(sumChan chan squareSum, wg *sync.WaitGroup) {

	sumStruct := <-sumChan
	fmt.Printf("Sum of square of %v is %v \n", sumStruct.number, sumStruct.sum)

	wg.Done()
}

func execute() {
	var wg sync.WaitGroup
	digitsChan := make(chan digits, 1000)
	squaresChan := make(chan digits, 1000)
	sumChan := make(chan squareSum, 1000)

	for i := 1; i <= 1000; i++ {
		wg.Add(4)
		
		go getDigits(i, digitsChan, &wg)
		go getSquare(digitsChan, squaresChan, &wg)
		go getSum(squaresChan, sumChan, &wg)
		go printSum(sumChan, &wg)

	}
	wg.Wait()
}

func main() {
	execute()
}
