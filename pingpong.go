package main

import (
    "fmt"
    "time"
)

func ping(ch chan string) {
    for {
        ch <- "ping"
        time.Sleep(time.Second) 
    }
}

func pong(ch chan string) {
    for {
        ch <- "pong"
        time.Sleep(time.Second)
    }
}

func main() {
    ch := make(chan string)

    go ping(ch)
    go pong(ch)

    for i := 0; i < 10; i++ { 
        msg := <-ch
        fmt.Println(msg)
    }
}
