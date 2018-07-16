package main

import (
	"fmt"
	"reflect"
)

// Boolean
//var isActive bool                 // 包级变量声明。
//var enable, disable = true, false // 包级忽略类型声明，从明确赋值中自动识别。
//
//isBool() {
//	var available bool // 局部一般声明
//	valid := false     // 局部简短声明
//	available = true   // 局部赋值操作
//
//	fmt.Println(reflect.TypeOf(available))
//	fmt.Println(reflect.TypeOf(valid))
//	fmt.Println(reflect.TypeOf(isActive))
//	fmt.Println(reflect.TypeOf(enable))
//	fmt.Println(reflect.TypeOf(disable))
//}
//

// int
func isInt() {
	var i int
	var a int8
	var b int16

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

	fmt.Println(i)
	fmt.Println(&i)
	i = 100
	fmt.Println(&i)
}

// string
func isString() {
	var frenchHello string                 // 声明变量为字符串的一般方式。
	var emptyString string = ""            // 声明了一个字符串变量，初始化为空字符串。
	no, yes, maybe := "no", "yes", "maybe" // 简短初始化多个变量。
	frenchHello = "Bonjour"                // 对便变量进行赋值。

    fmt.Println(frenchHello, emptyString, no, yes, maybe)

//    s := "hello"
//    c := []byte(s)     // 将字符串 s 转换为 []byte 类型。
//    c[0] = 'c'
//    s2 := string(c)    // 将 []byte 转换为 string 类型。
//    fmt.Println(s2)

//    s := "hello,"
//    m := " world"
//    a := s + m
//    fmt.Println(a)

    s := "hello"
    fmt.Println(&s)
    s = "c" + s[1:]  // 字符串虽不能更改，但是可以进行切片操作。
    fmt.Println(&s)
    fmt.Println(s)

//    z := `hello
//          world`
//    fmt.Println(z)

}


func main() {
    isString()
}
