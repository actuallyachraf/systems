package main

import (
	"fmt"
	"sync"
)

// Let's take example 2 a bit further

// In this example we want to concurrently print the following values.
// Thinking about sequentially you might think we'll see the messages printed
// in a certain order.
// Except that what you end up seeing is just Bonjour
func badConcurrentLoop() {
	var wg sync.WaitGroup

	for _, HiMsg := range []string{"Hello", "Hi", "Konichiwa", "Bonjour"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(HiMsg)
		}()
	}

	wg.Wait()
}

// Interestingly, HiMsg will fall out of scope.
// Go's runtime solves this by transfering the next value
// to the heap since it's still being used, except it will be
// stuck with the last value.

// Solving this requires forcing the HiMsg to go within scope.
// Here the first message to print will be Bonjour then Go will print Hello, Hi and Konichiwa
// respectively.
func GoodConcurrentLoop() {

	var wg sync.WaitGroup
	for _, HiMsg := range []string{"Hello", "Hi", "Konichiwa", "Bonjour"} {
		wg.Add(1)
		go func(HiMsg string) {
			defer wg.Done()
			fmt.Println(HiMsg)
		}(HiMsg)
	}

	wg.Wait()
}

func main() {
	badConcurrentLoop()
	GoodConcurrentLoop()
}
