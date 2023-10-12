package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/brunomvsouza/singleflight"
)

func main() {
	var (
		group singleflight.Group[string, string]
		wg    sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		res1, err1, shared1 := group.Do("key", func() (string, error) {
			time.Sleep(10 * time.Millisecond)
			return "func 1", nil
		})
		fmt.Println("res1:", res1)
		fmt.Println("err1:", err1)
		fmt.Println("shared1:", shared1)
	}()

	go func() {
		defer wg.Done()
		res2, err2, shared2 := group.Do("key", func() (string, error) {
			time.Sleep(10 * time.Millisecond)
			return "func 2", nil
		})
		fmt.Println("res2:", res2)
		fmt.Println("err2:", err2)
		fmt.Println("shared2:", shared2)
	}()

	wg.Wait()

	// Output:
	// res1: func 2
	// err1: <nil>
	// res2: func 2
	// err2: <nil>
	// shared2: true
	// shared1: true
}
