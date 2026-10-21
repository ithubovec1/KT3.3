package main

import (
	"fmt"
	"sync"
)

func sumPart(nums []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, n := range nums {
		sum += n
	}
	ch <- sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mid := len(numbers) / 2
	part1 := numbers[:mid]
	part2 := numbers[mid:]

	ch := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go sumPart(part1, ch, &wg)
	go sumPart(part2, ch, &wg)

	wg.Wait()
	close(ch)

	totalSum := 0
	for partialSum := range ch {
		totalSum += partialSum
	}

	fmt.Printf("Sum: %d\n", totalSum)
}