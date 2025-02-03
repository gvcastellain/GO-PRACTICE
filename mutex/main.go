package main

import (
	"fmt"
	"sync"
)

var value = 0
var mutex sync.Mutex

func incrementar(wg *sync.WaitGroup) {
	mutex.Lock()

	defer wg.Done()

	defer mutex.Unlock() //release acess first

	value++
	fmt.Println("value:", value)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementar(&wg)
	}

	wg.Wait()
	fmt.Println("final value:", value)
}
