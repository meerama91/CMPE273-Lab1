package main

import (
        "time"
        "fmt"
)

func sleeps(x int) {
    <-time.After(time.Second * time.Duration(x))
}

func main() {
    t1 := time.Now().Unix()
    sleeps(3)
    t2 := time.Now().Unix()

    fmt.Println("t1: ", t1)
    fmt.Println("t2: ", t2)
}