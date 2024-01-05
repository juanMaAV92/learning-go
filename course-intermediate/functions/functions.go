package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	c := make(chan int)
	go func() {
		fmt.Println("Hello")
		time.Sleep(5 * time.Second)
		fmt.Println("World")
		c <- 1
	}()

	<-c
}
