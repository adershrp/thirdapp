package main

import (
    "fmt"
    "runtime"
    "sync"
)

var counter = 0
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

// this is also a goroutine
func main() {
    /*
       By default the number of threads == nos cpu cores. this can be set to 1.
       Below line will give the default nos of threads or nos of cpu cores.
       runtime.GOMAXPROCS(-1) will return the current set of #threads
    */
    fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(-1))
    runtime.GOMAXPROCS(10) // setting the max number of threads
    fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(-1))
    
    for i := 0; i < 10; i++ {
        wg.Add(2)        // telling how many additional go routines
        m.RLock()        // read mutex called
        go sayHello()    // spin a new goroutine
        m.Lock()         // write mutex called
        go incrementer() // spin a new goroutine
    }
    
    /* asking main program to wait */
    wg.Wait()
}

// sayHello
func sayHello() {
    fmt.Println("Hello from Go Routine", counter)
    m.RUnlock() // read mutex released
    wg.Done()   // telling i'm done.
}

// incrementer
func incrementer() {
    counter++
    m.Unlock() // write mutex released
    wg.Done()
}
