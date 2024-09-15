package tests

import (
	"testing"
	
	"github.com/SantiSite/workerhttp/functions"
)

func TestFibonacci(t *testing.T) {
	t.Run("Fib 5", func(t *testing.T) {
		fib := functions.Fibonacci(5)
		if fib != 5 {
			t.Errorf("Fibonacci was incorret, got %d, expected %d", fib, 5)
		}
	})

	t.Run("Fib 8", func(t *testing.T) {
		fib := functions.Fibonacci(8)
		if fib != 21 {
			t.Errorf("Fibonacci was incorret, got %d, expected %d", fib, 21)
		}
	})
}
