package main

import (
    "fmt"
)

func slice1() {
    //slice := []byte{'a', 'b', 'c'}         // 初始化一个slice，它的值为 a,b,c
    var ar = [10]byte{'A', 'b', 'c', 'd', 'f', 'g', 'h', 'i', 'j'}  // 声明一个含有10个元素元素类型为byte的数组
    var a, b []byte  // 声明两个含有byte的slice

    a = ar[2:5]
    b = ar[3:5]

    fmt.Println("ar:", ar)
    fmt.Println("&ar:", &ar)
    fmt.Println("a:", a)
    fmt.Println("&a:", &a)
    fmt.Println("b:", &(b)[0])
    fmt.Println("&b:", &b)

}


func main () {
   slice1()
}


