package main

import (
    "fmt"
    "math"
    "math/rand"
)

func main() {
    fmt.Println("Now you have %g problems.",  math.Sqrt(7))
    fmt.Println("Now you have %d problems.",  math.Sqrt(7))
    fmt.Println("Now you have %g problems.",  rand.Intn(7))
}
