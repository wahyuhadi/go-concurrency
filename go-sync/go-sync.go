package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("Test")
		wg.Done()
	}()
	wg.Wait()
}


func count (num string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, num)
		time.Sleep(time.Millisecond * 500)
	}
}

