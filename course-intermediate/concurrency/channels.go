package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		ch <- 1
		wg.Add(1)
		go doSomething(i, &wg, ch)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Started gorutine %d\n", i)
	time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
	fmt.Printf("Finished gorutine %d\n", i)
	<-c
}
