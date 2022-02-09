package main

import "fmt"

func main() {
	a := 1			// 定义一个变量a，并赋值为1
	b := &a			// 通过取地址符取出变量a在内存中的地址
	fmt.Println(b)	// 打印取出的地址

	c := *b
	fmt.Println("c的值:", c)


	// 定义指针, 使用new进行初始化
	var point = new(int)
	*point = 1
	fmt.Printf("a的值：%d \n", *point)
	fmt.Printf("a的类型：%T \n", point)

}