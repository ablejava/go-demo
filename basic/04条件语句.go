package main

import "fmt"

func main() {
	score := 90
	if score >= 90 {
		fmt.Println("优秀")
	}

	if score > 90 {
		fmt.Println("优秀")
	} else {
		fmt.Println("一般")
	}

	if score > 90 {
		fmt.Println("优秀")
	} else if score == 90 {
		fmt.Println("及格")
	} else {
		fmt.Println("一般")
	}


	// switch
	switch {
	case score >= 90:fmt.Println("优秀")
	case score >= 80:fmt.Println("良好")
	case score >= 70:fmt.Println("一般")
	default:fmt.Println("默认")
	}
	// 在每个case后面都增加了fallthrough，则当case score >= 80匹配时，除了执行该case中的语句外，还会直接之后它后面的所有case，并且不会判断case是否为true。
	switch {
	case score >= 90:
		fmt.Println("优秀")
		fallthrough
	case score >= 80:
		fmt.Println("良好")
		fallthrough
	case score >= 70:
		fmt.Println("一般")
		fallthrough
	default:fmt.Println("默认")
	}

}