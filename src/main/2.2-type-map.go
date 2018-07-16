package main

import (
    "fmt"
)

func map1() {
    // 声明一个key是字符串，value为int的字典，这种方式的声明需要在使用之前使用make初始化
    // var numbers map[string]int

    // 初始化方式
    numbers := make(map[string]int)
    numbers["one"] = 1
    numbers["ten"] = 10
    numbers["three"] = 3

    fmt.Println(numbers["three"])
    fmt.Println(len(numbers))
}

func map2() {
    rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
    csharpRating, ok := rating["C"]
    if ok {
        fmt.Println(csharpRating)
    } else {
        fmt.Println("not key in rating")
    }

    fmt.Println(&rating)
    delete(rating, "C")
    fmt.Println(rating)
}

func map3() {
    m := make(map[string]string)
    m["Hello"] = "Bonjour"
    m1 := m
    m1["Hello"] = "Salut" // 现在m["hello"]的值已经是Salut了
    fmt.Println(m)
}

func main() {
    map3()
}
