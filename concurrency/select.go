package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := d1 / 2

	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)
		case channelMsh2 := <-c2:
			fmt.Println(channelMsh2)
		}
	}
}

func doSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
