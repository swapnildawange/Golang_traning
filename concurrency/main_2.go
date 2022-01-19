package main

import (
	"fmt"
	"sync"
)

var sum int = 0

func calculateSum(num chan int, mx *sync.Mutex, i int, wg *sync.WaitGroup) {
	v := <-num
	fmt.Println(v, i)
	defer wg.Done()
	// for j := range num {
	// 	fmt.Printf("The Value of %d is being processed by  %d goroutine \n", j, i)
	// 	mx.Lock()
	// 	sum += j * j
	// 	mx.Unlock()
	// }
}

func execute() {
	var mx sync.Mutex
	var wg sync.WaitGroup
	num := make(chan int, 1000)

	for count := 1; count <= 1000; count++ {
		num <- count
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go calculateSum(num, &mx, i, &wg)

	}

	close(num)
	wg.Wait()
	fmt.Println("The sum of squares is ", sum)
}

func main() {
	execute()
}
