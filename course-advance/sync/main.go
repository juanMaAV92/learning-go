package main

import "sync"

var (
	balance int = 100
)

// func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	defer lock.Unlock()

	lock.Lock()
	b := balance
	balance = b + amount

}

// func Balance(lock *sync.Mutex) int {
func Balance(lock *sync.RWMutex) int {
	defer lock.RUnlock()
	lock.RLock()
	b := balance
	return b
}

func main() {
	var wg sync.WaitGroup
	// var lock sync.Mutex
	var lock sync.RWMutex

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go Deposit(100*i, &wg, &lock)
	}
	wg.Wait()
	println(Balance(&lock))
}
