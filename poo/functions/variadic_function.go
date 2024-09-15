package main

import "fmt"

func sum(values ...int) (total int) {
	total = 0
	for _, value := range values {
		total += value
	}
	return
}

func printName(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValue(x int) (double, triple, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
	printName("Santi", "Juan")
	fmt.Println(getValue(2))
}
