package main

import "fmt"

//定义变量 声明常量 调用方法 数据类型

//大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
//大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。

func main() {
	fmt.Println("00000")
	aaa()
	bbb()
	arrayList()
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
  world` //` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出
	if !(a2 == b2) {
		fmt.Println(c2)
		fmt.Println(name + name2)
	}
}

func bbb() {
	s := "hello"
	s = "c" + s[2:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)
}

func arrayList() {
	//数组
	var arr1 [10]int

	arr1[1] = 1
	arr1[2] = 2

	fmt.Println(arr1[1])

	array1 := [3]int{1, 2, 3}
	array2 := [...]int{2, 2, 2, 2, 2}
	array3 := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	//在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice
	array4 := []int{1, 2, 3}
	fmt.Println(array4)
	//获取数组长度
	fmt.Println(len(array4))
	//数组新增一个元素
	array5 := append(array4, 3)
	fmt.Println(array5)
}
