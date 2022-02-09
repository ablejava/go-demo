package main

import (
	"fmt"
	"time"
)
//ch1 := make(chan int)			// 初始化一个int类型的通道
//ch2 := make(chan string, 2)		// 初始化一个string类型的通道，且缓冲区大小为2


func main() {
	//无缓冲通道在发送值的时候必须要有一个对应的 goroutine 接收值才行，**否则就会报错
	ch := make(chan int)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second)
}

func send(ch chan int) {
	//ch <- 10
	//fmt.Println("已发送数字 10 到通道中")

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

}

func receive(ch chan int) {
	//x := <- ch
	//fmt.Println("从通道接收到元素：", x)

	for  {
		x, ok := <- ch
		if ok {
			fmt.Println("接收到的值：", x)
		} else {
			break
		}
	}
}
// ------------------for range循环接收的方式判断发送方是否已经关闭通道，如果发送方关闭通道了，那么循环就会自动退出。------------
func rangeReceive(ch chan int) {
	for i := range ch {
		fmt.Println("接收到的值：", i)
	}
}

// ------------------------有缓存区通道可以没有接收channel -----------------------
func cacheChannel() {
	ch := make(chan int, 1)		// 初始化一个缓冲区为 1 的有缓冲通道
	ch <- 10					// 发送一个值到通道中
	fmt.Println("成功发送 10 到通道中")
}








// ------------------------单向通道------------------------

//chan <- 类型：只能发送值到通道中，不能从通道中接收值
//<- chan 类型：只能从通道中接收值，不能发送值到通道中
func oneWayChannel() {
	ch := make(chan int, 2)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second)
}

func oneWaySend(ch chan <- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func oneWayReceive(ch <- chan int) {
	for i := range ch {
		fmt.Println("接收到的值：", i)
	}
}
