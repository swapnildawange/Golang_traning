package main

import (
	"fmt"
	"sync"
)

var sum int = 0

func calculateSum(start chan int, routineIdx int, mx *sync.Mutex, wg *sync.WaitGroup) {

	defer wg.Done()

	startVal := <-start

	i := startVal

	for ; i < startVal+200; i++ {
		fmt.Printf("The Value of %d is being processed by  %d goroutine \n", i, routineIdx)
		mx.Lock()
		sum += i * i
		mx.Unlock()
	}
	start <- i
}

func execute() {
	var mx sync.Mutex
	var wg sync.WaitGroup

	wg.Add(5)

	start := make(chan int, 1000)
	start <- 0

	for i := 0; i < 5; i++ {
		go calculateSum(start, i, &mx, &wg)
	}

	wg.Wait()
	fmt.Println("The sum of squares is ", sum)
}

func main() {
	execute()
}
