package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Empezo lo gorutina principal.")
	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething1(i, &wg)
	}
	fmt.Println("Terminal la gorutina principal.")
}

func doSomething1(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started. %d.\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End.")
}
