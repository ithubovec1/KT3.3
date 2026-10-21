package main

import "fmt"

func generator(out chan<- int) {
	for i := 1; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func doubler(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * 2
	}
	close(out)
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go generator(ch1)
	go doubler(ch1, ch2)

	for res := range ch2 {
		fmt.Println(res)
	}
}