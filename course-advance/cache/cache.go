package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {
	fmt.Printf("Computing fib(%d)\n", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {
	s.Lock.RLock()
	exists := s.InProgress[job]

	if exists {
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Printf("Job %d is already in progress, waiting for result\n", job)
		fib := <-response
		fmt.Printf("Job %d is done, result: %d\n", job, fib)
		return
	}
	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Job %d is not in progress, starting it\n", job)
	fibo := ExpensiveFibonacci(job)

	s.Lock.RLock()
	pendingWorkers, exists := s.IsPending[job]
	s.Lock.RUnlock()

	if exists {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- fibo
		}
		fmt.Printf("Result sent to %d pending workers job:%d\n", len(pendingWorkers), job)
	}
	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, job := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(job)
	}
	wg.Wait()

}
