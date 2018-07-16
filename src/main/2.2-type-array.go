package main

import (
	"fmt"
)

func isArray1() {
	var arr [10]int                                 // 声明了一个int类型的数组
	arr[0] = 42                                     // 数组下标是从0开始的
	arr[1] = 13                                     // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0]) // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9])  // 返回未赋值的最后一个元素，默认返回0
}

func isArray2() {
	a := [3]int{1, 2, 3}   // 初始化一个长度为3的int数组（确定赋值初始化）
	b := [10]int{1, 2, 3}  // 初始化一个长度为10的int数组，其中前三个元素的值为1,2,3 其他默认为0
	c := [...]int{4, 5, 6} // 可以省略长度而采用 ... 的方式，Go会自动根据元素个数来计算长度
	fmt.Println(a, b, c)
}

// Double Array
func isArray3() {
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	edoubleArray := [2][4]int{{1, 2, 3, 4}, {5, 4, 7, 8}}
    fmt.Println(doubleArray,edoubleArray)
}

func main() {
	isArray3()
}
