package main

import (
    "fmt"
    "time"
)

const (
    logInfo    = "INFO"
    logWarning = "WARNING"
    logError   = "ERROR"
)

type logEntry struct {
    time     time.Time
    severity string
    message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func main() {
    go getLogger()
    logCh <- logEntry{time: time.Now(), severity: logInfo, message: "App Started!"}
    for i := 0; i < 10; i++ {
        logCh <- logEntry{time: time.Now(), severity: logInfo, message: fmt.Sprint("Message", i)}
        time.Sleep(100 * time.Millisecond)
    }
    logCh <- logEntry{time: time.Now(), severity: logInfo, message: "App Shutdown!"}
    time.Sleep(100 * time.Millisecond)
    doneCh <- struct{}{}
}

func getLogger() {
    for {
        select {
        case entry := <-logCh:
            fmt.Println(entry.time.Format(time.RFC3339), entry.severity, entry.message)
        case <-doneCh:
            fmt.Println("closing the channel")
            break
            // default:
            //     fmt.Println("Default Case")
        }
    }
}
