package main

import (
	"fmt"
	"sync"
)

// Example of Data Race :
// Given that value is a shared resource we effectively have a race condition
// We can denote the sections of the condition as *critical*
// So in the snippet below we have 3 critical sections
// 1. The go routine itself
// 2. The if clause check
// 3. The else clause check
/*
	var value int
	go func() {value++}
	if value == 0 {
		fmt.Println("value is equal to 0")
	} else {
		fmt.Println("value is equal to ",value)
	}

*/

// In order to fix this using primitive concurrency methods we introduce
// locks, locks are implemented as a *Mutex*
// We still have an issue here, execution order
// we don't know whether the go routine will run first or whether
// the if/else clause .
// The order of operations is still non-deterministic
func main() {

	var lock sync.Mutex
	var value int
	// lock the mutex within the go routine
	go func() {
		lock.Lock()
		value++
		lock.Unlock()
	}()

	// lock the mutex to check the conditional clauses
	// locking is similar to saying "you have exclusive access to this"
	lock.Lock()
	if value == 0 {
		fmt.Println("value is equal to ", value)
	} else {
		fmt.Println("value is equal to ", value)
	}
	lock.Unlock()
}
