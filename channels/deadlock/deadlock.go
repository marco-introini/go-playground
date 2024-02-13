package main

import "fmt"

func main() {
	dataChan := make(chan int)
	for i := 0; i < 1000; i++ {
		dataChan <- i
	}

	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}
}
