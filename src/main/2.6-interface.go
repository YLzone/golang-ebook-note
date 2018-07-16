package main

import (
    "fmt"
//  "reflect"
)

type I interface {
    Get()
    Set()
}

type S struct {
    Name string
}

func (s *S) Get() {
    fmt.Println(s.Name) 
}

func (s *S) Set() {
    s.Name = "set"
}

func F() {
    fmt.Println("is a function")
}


func main() {
    var i I

    s := S{"xiaoming"} 
    i = &s

    fmt.Println(s.Name)
    i.Set()
    i.Get()
    fmt.Println(s.Name)

}
