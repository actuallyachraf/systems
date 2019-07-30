package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("I am the first go routine, I am working (actually sleeping) !")
		time.Sleep(1)
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("I am the second go routine I am actually sleeping")
		time.Sleep(1)
	}()

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("I am the third go routine brother of the fourth")
		time.Sleep(1)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("I am the fourth go routine brother of the third")
		time.Sleep(1)
	}()
	// wait for all go routines to finish
	wg.Wait()

	hello := func(wg *sync.WaitGroup, uid int) {
		defer wg.Done()
		fmt.Println("Hello from ", uid)
	}

	const numberOfHelloers = 8

	wg.Add(numberOfHelloers)
	for i := 0; i < numberOfHelloers; i++ {
		go hello(&wg, i+1)
	}

	wg.Wait()
}
