package main

import "fmt"

func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		//fmt.Printf("Generator %d.\n", i)
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		//fmt.Printf("Double %d.\n", value*2)
		out <- 2 * value
	}
	close(out)
}

func Print(out <-chan int) {
	for value := range out {
		//fmt.Printf("Print %d.\n", value)
		fmt.Println(value)
	}
}

func main() {
	c := make(chan int)
	out := make(chan int)

	go Generator(c)
	go Double(c, out)
	Print(out)
}
