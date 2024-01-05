package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("S", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	m := make(map[string]int)
	m["key"] = 19
	fmt.Println(m["key"])

	s := []int{5, 6, 7}
	for index, value := range s {
		fmt.Printf("El valor del index %d es: %d\n", index, value)
	}
	s = append(s, 13)
	for index, value := range s {
		fmt.Printf("El valor del index %d es: %d\n", index, value)
	}

	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	fmt.Println("Doing something")
	time.Sleep(1 * time.Second)
	fmt.Println("Done doing something")
	c <- 1
}
