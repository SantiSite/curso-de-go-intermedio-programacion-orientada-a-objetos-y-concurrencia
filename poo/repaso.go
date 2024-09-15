package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int // definir variables de forma explicita
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	myValue, err := strconv.ParseInt("7", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// maps
	m := make(map[string]string)
	m["Colombia"] = "Bogota"
	fmt.Printf("%+v\n", m)
	fmt.Printf("%v\n", m["Colombia"])

	// slice
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Printf("indice: %d, valor: %d\n", index, value)
	}
	s = append(s, 4)
	fmt.Println(s[3])

	// Gorrutinas
	c := make(chan int)
	go doSomething(c)
	<-c

	// apuntadores
	g := 25
	h := &g
	fmt.Println("direccion de memoria de g: ", h)
	fmt.Println("valor:", *h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
