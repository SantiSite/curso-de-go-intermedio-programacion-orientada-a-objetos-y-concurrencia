package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Este canal va a transmitir y obtener datos de tipo int con un tamaño de dos datos
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// Lo que hacemos es darle un cupo a la gorutina que esta en la linea 18, por lo tanto, se ejecutan de dos en dos las gorutinas
		c <- 1
		wg.Add(1)
		go doSomething2(i, &wg, c)
	}
	wg.Wait()
}

func doSomething2(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Id: %d. Started.\n", i)
	time.Sleep(4 * time.Second)

	fmt.Printf("Id: %d. Finished\n", i)
	// Liberamos el cupo cuando se lee el dato que viene en el canal
	<-c
}
