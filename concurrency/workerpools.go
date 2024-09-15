package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 7, 8, 9, 10, 12, 30, 38, 39, 40, 41, 42, 43, 44}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	fmt.Println("Se empieza a crear los trabajadores:")
	for i := 0; i < nWorkers; i++ {
		fmt.Printf("Trabajador con id %d creado\n", i)
		go Worker(i, jobs, results)
	}
	fmt.Printf("\n")

	fmt.Println("Se empieza a asignar las tareas en el canal de jobs:")
	for _, value := range tasks {
		fmt.Printf("Se asigna la tarea %d\n", value)
		jobs <- value
	}
	close(jobs)
	fmt.Printf("\n")

	for i := 0; i < len(tasks); i++ {
		<-results
	}

	fmt.Printf("Fin de programa!\n")
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id: %d. Started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id: %d, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
