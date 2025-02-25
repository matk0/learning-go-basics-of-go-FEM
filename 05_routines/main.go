package main

import "fmt"
import "time"

func printMessage(text string, channel chan string) {
    for i := 0; i < 5; i++ {
        fmt.Println(text)
        time.Sleep(500 * time.Millisecond)
    }
    channel <- "Done!"
}

func main() {
    channel := make(chan string)
    go printMessage("Go is great!", channel)
    response := <- channel
    fmt.Println(response)
}
