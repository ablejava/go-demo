package main

import "fmt"

func declareMap() {
	// 当前声明的m为nil map，无法直接使用,不会分配内存空间
	var m map[string]int
	// 使用make初始化之后可正常使用
	m = make(map[string]int)
	// 朝m里面放一个键为one，值为1的键值对数据
	m["one"] = 1
	m["two"] = 2
	// 取出键为one对应的值
	fmt.Println(m["one"])
	fmt.Println(m["two"])
}

func init() {
	m := map[string]int{"one" : 1, "two" : 2, "three" : 3}
	fmt.Println(m)
}

func forMap() {
	m := map[string]int{"one" : 1, "two" : 2}
	for key := range m {
		fmt.Printf("键%s对应的值为%d \n", key, m[key])
	}
}

func existMap() {
	m := map[string]int{"zero" : 0, "one" : 1, "two" : 2}
	value, ok := m["zero"]
	if ok {
		fmt.Println("zero存在，value：", value)
	} else {
		fmt.Println("zero不存在，value：", value)
	}
	value, ok = m["three"]
	if ok {
		fmt.Println("three存在，value：", value)
	} else {
		fmt.Println("three不存在，value：", value)
	}
}

func delMap() {
	m := map[string]int{"zero" : 0, "one" : 1, "two" : 2}
	// 删除m中key为two的元素
	delete(m, "two")
	value, ok := m["two"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("two不存在")
	}
}