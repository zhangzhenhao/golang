package main

import (
    "fmt"
)

func send(data int, channel chan int) {
    fmt.Println("send %v to channle", data)
    channel <- data
}

func read(channel chan int) {
    d := <- channel
    fmt.Println("read from c, and the data is:%v", d)
}

func main() {

    c := make(chan int,2)
    send(1, c)
    send(2, c)
    go send(3,c)
    read(c)
    read(c)
    read(c)
    go read(c)
    send(2, c)
    fmt.Println("success")
}
