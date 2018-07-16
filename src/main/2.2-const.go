package main

import (
	"fmt"
	"reflect"
)

const constName string = "Is Const"
const Pi float32 = 3.1415926
const i = 9999
const MaxThread = 10
const prefix = "astaxie_"

func main() {
	fmt.Println(constName)
	fmt.Println("Const MaxThred Type:", reflect.TypeOf(MaxThread))
	fmt.Println("Const Pi Type:", reflect.TypeOf(Pi))
}
