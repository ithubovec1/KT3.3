package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Num  int
	Fact uint64
}

func worker(id int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		var fact uint64 = 1
		for i := 1; i <= n; i++ {
			fact *= uint64(i)
		}
		results <- Result{Num: n, Fact: fact}
	}
}

func main() {
	numJobs := 20
	jobs := make(chan int, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Printf("Факториал числа %d равен %d\n", res.Num, res.Fact)
	}
}