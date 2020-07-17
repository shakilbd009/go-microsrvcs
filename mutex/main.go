package main

import (
	"fmt"
	"sync"
)

var (
	counter = 0
	mu      sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go goUpdateCounter(&wg)
	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("final counter: %v", counter))
}

func goUpdateCounter(wg *sync.WaitGroup) {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}
