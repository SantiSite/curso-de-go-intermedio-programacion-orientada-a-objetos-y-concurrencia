package main

import "fmt"

func main() {
	// buffered channel
	c := make(chan int, 1)
	c <- 1
	fmt.Println("c:", <-c)

	// unbuffered channel
	c2 := make(chan string)
	defer func() {
		close(c)
		close(c2)
	}()

	go goFunction(c2)
	fmt.Println("c2:", <-c2)
}

func goFunction(c chan string) {
	c <- "SantiSite"
}
