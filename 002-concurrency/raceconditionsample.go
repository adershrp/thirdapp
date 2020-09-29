package main

import (
    "fmt"
    "time"
)

/**
To Analyse race condition run program with go run -race
*/
func main() {
    var msg = "Hello"
    go func() {
        fmt.Println(msg)
    }()
    msg = "GoodBye"
    time.Sleep(time.Millisecond * 100)
}
