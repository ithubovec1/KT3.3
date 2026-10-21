package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	words := strings.Fields(text)

	var wg sync.WaitGroup
	var mu sync.Mutex
	count := 0

	for _, word := range words {
		if word == "" {
			continue
		}
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}(word)
	}

	wg.Wait()
	fmt.Printf("Количество слов: %d\n", count)
}