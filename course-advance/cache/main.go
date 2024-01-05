package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock  sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	res, ok := m.cache[key]
	m.lock.Unlock()
	if !ok {
		m.lock.Lock()
		res.value, res.err = m.f(key)
		m.cache[key] = res
		m.lock.Unlock()
	}
	return res.value, res.err
}

func GetFibonacci(n int) (interface{}, error) {
	if n < 0 {
		return 0, errors.New("invalid input")
	}
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	fibo := []int{1, 2, 4, 5, 8, 9, 10, 38, 42, 40}
	var wg sync.WaitGroup

	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Printf("Error: %v\n", err)
			}
			fmt.Printf("Fibonacci(%d) = %d, err = %v, time = %v\n", index, value, err, time.Since(start))
		}(n)
	}

	wg.Wait()

	// realizar un delay de 5 segundos para que se pueda ver el efecto de la cache
	time.Sleep(5 * time.Second)
	fmt.Println("+--------------------------+")

	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Printf("Error: %v\n", err)
			}
			fmt.Printf("Fibonacci(%d) = %d, err = %v, time = %v\n", index, value, err, time.Since(start))
		}(n)
	}

	wg.Wait()
}
