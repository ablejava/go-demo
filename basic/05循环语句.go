package main

import "fmt"

func main() {
	// 计算从1到100的累加结果
	var sum1 int
	for i := 1; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println("1到100累加结果为：", sum1)


	// 第二种方法
	var sum2 int = 1
	i:= 1
	for i <= 100 {
		sum2 *= i
		i++
	}
	fmt.Println(sum2)

	// 第三中遍历数组或者集合
	numbers := [9]int {1,2,3,4,5,6,7,8,9}
	for i2, number := range numbers {
		fmt.Println("下标index:%s, 数值：%s", i2, number)
	}
}