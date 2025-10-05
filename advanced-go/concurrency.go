package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[string]int)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)

			// This might cause a race condition
			m[fmt.Sprintf("key-%d", i)] = i
		}(i)
	}
	wg.Wait()
	fmt.Println("Map:", m)

}
