package functions

// Fibonacci calcula el valor de la seria, en donde n es el valor de la posición a calcular.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
