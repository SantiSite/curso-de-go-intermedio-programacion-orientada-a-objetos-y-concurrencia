package examples_test

import (
	"fmt"

	"github.com/SantiSite/workerhttp/functions"
)

func ExampleFibonacci() {
	fib1 := functions.Fibonacci(5)
	fib2 := functions.Fibonacci(8)
	fmt.Println(fib1)
	fmt.Println(fib2)

	// Output:
	// 5
	// 21
}
