package main

import (
	"fmt"
	"sync"
)

// Given the following example
// func sayHi(){ fmt.Println("hello") }
// running it concurrently is as easy as
// go sayHi()
// The issue that stands is how you insure that the function
// will execute at all, Go will schedule the execution, but
// if your main routine exits before that what do you do ?
// A solution is to Sleep for certain time, except that
// it introduces data races.
// let's take a look at the synchronization using WaitGroups
// In this example waitgroup will deterministically block main
// until sayHello is executed, this guarrentes that your go routine
// is executed without worrying about data races.
func main() {

	var wg sync.WaitGroup

	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hi")
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()
}
