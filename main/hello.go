package main

import "fmt"

//定义变量 声明常量 调用方法 数据类型

func main() {
	fmt.Println("00000")
	aaa()
	bbb()
}

func aaa() {
	var a, b, c = 1, 2, 3
	isDone := true
	if isDone {
		fmt.Println(a + b + c)
		fmt.Println(a - b - c)
	}
	var a2, b2, c2 = "wo", "444", "999"
	const name = "ZhuYilin"
	var name2 = `hello
  world`
	if !(a2 == b2) {
		fmt.Println(c2)
		fmt.Println(name + name2)
	}
}

func bbb() {
	s := "hello"
	s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)
}
