package main

import (
    "fmt"
    "runtime"
    "sync"
)

var wg sync.WaitGroup

func main() {
    cpu := runtime.NumCPU()
    goarch := runtime.GOARCH
    goos := runtime.GOOS
    
    fmt.Printf("OS\t%v\n", goos)
    fmt.Printf("ARCH\t%v\n", goarch)
    fmt.Printf("CPU\t%v\n", cpu)
    fmt.Printf("Goroutine\t%v\n", runtime.NumGoroutine())
    wg.Add(2)
    go bar()
    go foo()
    fmt.Printf("Goroutine\t%v\n", runtime.NumGoroutine())
    wg.Wait()
}

func bar() {
    for i := 0; i < 10; i++ {
        fmt.Println("Bar", i)
    }
    wg.Done()
}
func foo() {
    for i := 0; i < 10; i++ {
        fmt.Println("Foo", i)
    }
    wg.Done()
}
