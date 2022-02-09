package main

import "fmt"

// 定义切片
var cutSlices []int
// make函数定义，length为初始化长度，capacity切片容量
// make([]int, length, capacity)

// 初始化
//s := arr[startIndex:endIndex]	// 从 startIndex 到 endIndex - 1 初始化为一个切片
//s := arr[startIndex:]		// 从 startIndex 到 数组结尾 初始化为一个切片
//s := arr[:endIndex]		// 从 数组开始 到 endIndex - 1 初始化为一个切片
//s := arr[:]			// 从 数组开始 到 数组结尾 初始化为一个切片（整个数组）
//
//s :=[] int {1,2,3}

func main() {
	arr := [5]int{1,2,3,4,5}
	s1 := arr[:]
	fmt.Println(s1)

	s2 := arr[2:]
	fmt.Println(s2)

	s3 := arr[:3]
	fmt.Println(s3)

	s4 := arr[1:4]
	fmt.Println(s4)

	// 定义一个切片, append追加一个元素
	s5 := []int{1,2,3}
	s5 = append(s5, 4)
	fmt.Println("append:", s5)

	// copy复制一个元素
	s6 := make([]int, 4)
	copy(s6, s5)
	fmt.Println("copy:", s6)


	// 截取s[开始索引:结束索引]
	s7 := []int{1,2,3,4,5}
	fmt.Println("s[1:3]截取：", s7[1:3])		// 截取索引1（包含）到索引3（不包含）
	fmt.Println("s[:4]截取：", s7[:4])			// 默认开始索引为0
	fmt.Println("s[2:]截取：", s7[2:])			// 默认结束索引为len


	// len 和 cap 方法
	s8 := make([]int, 3)
	fmt.Printf("len:%d, cap:%d \n", len(s8), cap(s8))		// 此时len=3，cap=3

	s8 = append(s8, 1)
	fmt.Printf("len:%d, cap:%d \n", len(s8), cap(s8))		// append一个元素，len=4，大于cap，所以底层数组扩容为cap *2,然后将之前的数组全部复制到新的数组中，cap=6

	s8 = append(s8, 1)
	fmt.Printf("len:%d, cap:%d \n", len(s8), cap(s8))		// append一个元素，len=5，小于cap，底层数组不变，cap=6

	s8 = append(s8, 1)
	fmt.Printf("len:%d, cap:%d \n", len(s8), cap(s8))		// append一个元素，len=6，小于cap，底层数组不变，cap=6

	s8 = append(s8, 1)
	fmt.Printf("len:%d, cap:%d \n", len(s8), cap(s8))		// append一个元素，len=7，大于cap，所以底层数组扩容，cap=12

}
