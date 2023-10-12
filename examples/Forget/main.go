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

	wg.Add(4)

	go func() {
		defer wg.Done()
		res1, err1, shared1 := group.Do("key", func() (string, error) {
			time.Sleep(1 * time.Second)
			return "func 1", nil
		})
		fmt.Println("res1:", res1)
		fmt.Println("err1:", err1)
		fmt.Println("shared1:", shared1)
	}()

	go func() {
		defer wg.Done()
		res2, err2, shared2 := group.Do("key", func() (string, error) {
			time.Sleep(1 * time.Second)
			return "func 2", nil
		})
		fmt.Println("res2:", res2)
		fmt.Println("err2:", err2)
		fmt.Println("shared2:", shared2)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		group.Forget("key")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		res3, err3, shared3 := group.Do("key", func() (string, error) {
			time.Sleep(1 * time.Second)
			return "func 3", nil
		})
		fmt.Println("res3:", res3)
		fmt.Println("err3:", err3)
		fmt.Println("shared3:", shared3)
	}()

	wg.Wait()

	// Output:
	// res2: func 1
	// err2: <nil>
	// shared2: true
	// res1: func 1
	// err1: <nil>
	// shared1: true
	// res3: func 3
	// err3: <nil>
	// shared3: false
}
