package main

import (
    "fmt"
)

func new1() {
    var i * int
    i = new(int)
    *i=10
    fmt.Println("new1:", *i)
}

func new2() {
    var i int
    var z *int
    z = &i
    *z = 20
    fmt.Println("new2:", *z)
    // fmt.Println("new2:", &*z)
    // fmt.Println("new2:", *&*z)
}

func main() {
    new1()
    new2()
}
