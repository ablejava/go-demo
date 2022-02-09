package main

import (
	"fmt"
	"time"
)

func blockSelect() {
	intChan := make(chan int)
	stringChan := make(chan string)

	go func(ch chan int) {
		time.Sleep(time.Second * 2)
		ch <- 1
	}(intChan)

	go func(ch chan string) {
		time.Sleep(time.Second)
		ch <- "a"
	}(stringChan)

	//select 会一直堵塞，直到其中一个 case 的 channel 接收到数据就执行其对应的输出语句，由于在往 intChan 中发送数据之前有两秒钟的睡眠时间， stringChan 只有一秒钟的睡眠时间，所以当前代码中永远都是 stringChan 先好，最后的结果则是在控制台打印字母 a
	select {
	case i := <- intChan:		// 监听 intChan 的操作
		fmt.Println(i)
	case s := <- stringChan:	// 监听 stringChan 的操作
		fmt.Println(s)
	}
}

func noBlockSelect() {
	intChan := make(chan int)
	stringChan := make(chan string)

	go func(ch chan int) {
		time.Sleep(time.Second * 2)
		ch <- 1
	}(intChan)

	go func(ch chan string) {
		time.Sleep(time.Second)
		ch <- "a"
	}(stringChan)

	//在 select 中增加 default 语句后则不会一直堵塞等待
	select {
	case i := <- intChan:		// 监听 intChan 的操作
		fmt.Println(i)
	case s := <- stringChan:	// 监听 stringChan 的操作
		fmt.Println(s)
	default:
		fmt.Println("default操作")
	}
}

func multipleChannel() {

	intChan := make(chan int)
	stringChan := make(chan string)

	go func(ch chan int) {
		ch <- 1
	}(intChan)

	go func(ch chan string) {
		ch <- "a"
	}(stringChan)

	time.Sleep(time.Second)		// 睡眠一秒钟，确保两个 channel 都准备好

	//intChan 和 stringChan 都是准备就绪状态，则这时候 select 会随机选择一个 case 分支执行
	select {
	case i := <- intChan:
		fmt.Println(i)
	case s := <-stringChan:
		fmt.Println(s)
	}

}