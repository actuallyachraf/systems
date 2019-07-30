# Sync Package Notes

## WaitGroup

WaitGroup is as it sounds a Waiting Group for Go routines.
It's useful when you want to throw a few go routines but don't care much
about their results.

```go

var wg sync.WaitGroup

wg.Add(1)

go func(){
    defer wg.Done()
    fmt.Println("I am the first go routine, I am working (actually sleeping) !")
    time.Sleep(1)
}()

wg.Add(1)

go func(){
    defer wg.Done()
    fmt.Println("I am the second go routine I am actually sleeping")
    time.Sleep(1)
}()

// add takes the number of go routines for example I can group two go routines
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
```

A more complex example would be to group the tracking of go routines.
```go


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

```

