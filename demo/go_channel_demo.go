package main

import (
    "fmt"
)


func sum(values []int , channel chan int) {
    sum := 0
    for _, v := range values {
        sum += v
    }
    channel <- sum

}


func main() {
    items :=  []int { 1, 2, 3, 4}
    var c chan int
    // chan must make it. Looks same with map
    c = make(chan int)
    go sum(items[:len(items)/2], c)
    go sum(items[len(items)/2:], c)
    x, y :=  <-c, <-c
    fmt.Println(x, y)
}
