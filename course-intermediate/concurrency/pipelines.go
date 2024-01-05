package main

func Generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func Square(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func Print(c <-chan int) {
	for i := range c {
		println(i)
	}
}

func main() {
	generator := make(chan int)
	square := make(chan int)

	go Generator(generator)
	go Square(generator, square)
	Print(square)

}
