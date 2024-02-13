package main

import (
	"fmt"
	"sync"
	"time"
)

func DoWork(i int) int {
	time.Sleep(3 * time.Second)
	return i
}

func main() {
	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}

		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork(i)
				dataChan <- result
			}()
		}

		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Println("n =", n)
	}

}
