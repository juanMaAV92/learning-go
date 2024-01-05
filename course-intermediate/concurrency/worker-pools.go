package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started FIB job %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker %d finished FIB job %d with a result of %d\n", id, job, fib)
		results <- Fibonacci(fib)
	}
}

func main() {
	tasks := []int{3, 4, 5, 30, 7, 8, 9, 10}

	numWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < numWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}
