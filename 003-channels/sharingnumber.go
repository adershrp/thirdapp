package main

import (
    "fmt"
    "sync"
)

var wg = sync.WaitGroup{}

func main() {
    
    ch := make(chan int)
    wg.Add(2)
    go func() {
        // data receiving from channel
        i := <-ch
        fmt.Println("Received from Channel", i)
        
        i = 56
        fmt.Println("Sending to channel", i)
        // data sending to channel
        ch <- i
        wg.Done()
    }()
    
    go func() {
        j := 96
        fmt.Println("Sending to channel", j)
        // data sending to channel
        ch <- j
        // data receiving from channel
        j = <-ch
        fmt.Println("Received from Channel", j)
        wg.Done()
    }()
    wg.Wait()
}
