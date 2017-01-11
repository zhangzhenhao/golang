package main

import (
    "fmt"
    "time"
)


func say(data string) {

    for i := 0; i < 5; i++ {
        time.Sleep(1* time.Second)
        fmt.Println("%v+%v",data,time.Now())
        fmt.Println("the i is %v", i)
    }
}


func main() {
    go say("hi")
    say("go")
    //fmt.Println("the Millisecond is %v", time.Millisecond)
    //time.Sleep(1000*time.Millisecond)
    //time.Sleep(2*time.Second)
} 
