package main

import "fmt"

func variable() {
	// 定义变量的方式
	var s string
	fmt.Println(s) // string 默认值空串
	var a, b int
	fmt.Println(a,b)
	fmt.Println("-----------------")

	// 使用 = 给变量赋值
	s = "hello world"
	fmt.Println(s)
	a,b=1,2
	fmt.Println(a,b)

	// 使用 := 给变量赋值，可以不定义变量类型
	n := 3
	fmt.Println(n)

	// 一次声明多个变量
	var(
		m string = "hello world"
		x int = 1
		z int = 2
	)
	fmt.Println(m, x, z)

	// 定义常量
	const (
		 i = ""
		 j = 1
	)
	fmt.Println(i)
	fmt.Println(j)
	const PI = 3.1415
	const E float32 = 2.7182
	fmt.Println(PI)
	fmt.Println(E)

}

func main() {
	variable()
}
