package main

import "fmt"

func main() {
	// 声明数组
	var intArr [5]int
	var strArr [5]string
	// 声明可变长数组
	var floatArr = [...]float32{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(floatArr)

	// 数组初始化
	intArr = [5]int{1,2,3,4,5}
	fmt.Println(intArr)
	strArr[0] = "a"
	strArr[1] = "b"
	strArr[2] = "c"
	strArr[3] = "d"
	strArr[4] = "e"

	for i, f := range floatArr {
		fmt.Println("index=%s, value=%s", i, f)
	}


	// 二维数组
	var intSecondArr [3][4]int
	intSecondArr = [3][4]int {
		{0,1,2,3},
		{4,5,6,7},
		{8,9,10,11},
	}
	fmt.Println(intSecondArr)
	// 输出二维数组
	for i, v := range intSecondArr {
		for j, v1 := range v {
			fmt.Printf("第%d行第%d列：%d \n", i + 1, j + 1, v1)
		}
	}

	for i := 0; i < len(intSecondArr); i++ {			// len(intSecondArr)获取到的是intSecondArr的行数，intSecondArr[i] 表示一个一维数组
		for j := 0; j < len(intSecondArr[i]); j++ {	// 按照一维数组的方式遍历即可
			fmt.Printf("第%d行第%d列：%d \n", i + 1, j + 1, intSecondArr[i][j])
		}
	}


}