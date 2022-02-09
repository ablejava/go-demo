package main

import "fmt"

// --------------返回值---------------------
// 声明一个函数，传入两个int类型整数，返回两个数相加的和
func add(n1 int, n2 int) int {
	return n1 + n2
}

// 声明一个函数，传入两个int类型整数，返回两个结果，一个是两数之和，一个是两数之差
func addAndSub(n1 int, n2 int) (add int, sub int) {
	add = n1 + n2
	sub = n1 - n2
	return add, sub
}

//-----------值传递和引用传递-----------------
// 定义一个函数，将传递进来的参数的值变为10
func change1(num int) {
	num = 10
}
// 定义一个函数，将传递进来的参数的值变为10
func change2(num *int) {
	*num = 10
}

// ---------------延迟调用------------------
func delay1() {
	fmt.Println("延迟调用...")
}
func delay2() {
	fmt.Println("延迟调用...")
}
func delay3() {
	fmt.Println("延迟调用...")
}


func main() {

	// 只接受add的值，不接收差的值
	add, _ := addAndSub(1,2)
	fmt.Println(add)



	// 定义局部变量
	num := 1
	fmt.Println("改变前num的值：", num)
	// 将实际参数传递给函数change
	change1(num)
	fmt.Println("改变后num的值：", num)

	// 将实际参数传递给函数change
	change2(&num)
	fmt.Println("改变后num的值：", num)


	// 延迟调用
	// 延迟调用会等待主函数都执行完成之后再执行。会按照delay3，delay2，delay1的顺序执行
	defer delay1()
	defer delay2()
	defer delay3() // 例如：数据库连接关闭
	fmt.Println("主函数执行。。。。。。。")
}