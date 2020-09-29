package main

import (
    "fmt"
    "sync"
)

var wg = sync.WaitGroup{}

func main() {
    // simpleOne()
    /**
      For creating a buffered channel pass extra integer param.
    */
    ch := make(chan int, 50)
    
    wg.Add(2)
    /**
      Channel receiving (<-) value, and the type is integer.
    */
    go func(ch <-chan int) {
        for {
            /**
              This way, on a infinite loop verify if the channel is open then get the data.
              If closed, break the loop.
            */
            if i, ok := <-ch; ok {
                fmt.Println("Received data from channel", i)
            } else {
                break
            }
        }
        wg.Done()
    }(ch)
    /**
      Integer value sending (<-) to channel.
    */
    go func(ch chan<- int) {
        for i := 0; i < 10; i++ {
            fmt.Println("Sending data to Channel", 1*i)
            ch <- 1 * i
            fmt.Println("Sending data to Channel", 2*i)
            ch <- 2 * i
        }
        /**
         *   Closing the channel, after this nothing can be send to this channel.
         */
        close(ch)
        wg.Done()
    }(ch)
    wg.Wait()
}

func simpleOne() {
    ch := make(chan int)
    // receiving channel
    for j := 0; j < 10; j++ {
        
        wg.Add(2)
        go func(ch <-chan int) {
            i := <-ch
            fmt.Println("Received data from channel", i)
            wg.Done()
        }(ch)
        
        go func(ch chan<- int, j int) {
            fmt.Println("Sending data to Channel", j)
            ch <- j
            wg.Done()
        }(ch, j)
    }
    wg.Wait()
}
