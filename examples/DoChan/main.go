package main

import (
	"fmt"

	"github.com/brunomvsouza/singleflight"
)

func main() {
	var group singleflight.Group[string, string]

	semaphore := make(chan struct{})

	res1c := group.DoChan("key", func() (string, error) {
		fmt.Printf("func 1 begin\n")
		defer fmt.Printf("func 1 end\n")
		<-semaphore
		return "func 1", nil
	})

	res2c := group.DoChan("key", func() (string, error) {
		fmt.Printf("func 2 begin\n")
		defer fmt.Printf("func 2 end\n")
		<-semaphore
		return "func 2", nil
	})

	close(semaphore)

	res1 := <-res1c
	res2 := <-res2c

	// Results are shared by functions executed with
	// duplicate keys.
	fmt.Println("Shared:", res2.Shared)
	// Only the first function is executed: it is registered and
	// started with "key", and doesn't complete before the second
	// funtion is registered with a duplicate key.
	fmt.Println("Equal results:", res1.Val == res2.Val)
	fmt.Println("Result:", res1.Val)

	// Output:
	// Shared: true
	// Equal results: true
	// Result: func 1
}
