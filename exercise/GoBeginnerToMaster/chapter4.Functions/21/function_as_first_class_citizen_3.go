package main

import "fmt"

// 我们想将 MyAdd 函数赋值给 BinaryAdder 接口，直接赋值是不行的，我们需要一个底层函数类型与 MyAdd 一致的自定义类型的显式转换，这个自定义类型就是 MyAdderFunc
// 该类型实现了 BinaryAdder 接口，这样在经过 MyAdderFunc 的显示类型转换后，MyAdder 被赋值给了 BinaryAdder的变量 i。
// 这样，通过 i 调用的 Add 方法实际上就是 MyAdd 函数
func main() {
	var i BinaryAdder = MyAdderFunc(MyAdd)
	fmt.Println(i.Add(5, 6))
}

type BinaryAdder interface {
	Add(int, int) int
}

type MyAdderFunc func(int, int) int

func (f MyAdderFunc) Add(x, y int) int {
	return f(x, y)
}

func MyAdd(x, y int) int {
	return x + y
}
