package main

import (
    "fmt"
)

func  fibonacci(channel, quite_channel chan int) {
    x, y := 0, 1
    for {
        select {
            case channel <- x: 
                x, y = y, x + y 
            case <-quite_channel:
                fmt.Println("quite.")
                return
        }
    }
}

func main() {
    c, q :=  make(chan int), make(chan int)
    // unname func
    go func () {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        q <- 0
    }()
    fibonacci(c, q)

}
