package main

import (
	"fmt"
	"runtime"
	"time"
)

func testGoroutine() {
	fmt.Println("test goroutine...")
}
func testGoroutine1(i int) {
	fmt.Println("test goroutine...", i)
}

func main() {
	go testGoroutine()
	fmt.Println("test main...")

	for i := 0; i < 10; i++ {
		go testGoroutine1(i)
	}
	// 参数是 time.Second
	time.Sleep(3)
}

// -------------------当前goroutine让出cpu------------------------
func gosched() {
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("test")
		}
	}()
	runtime.Gosched()		// main 协程让出CPU，供子协程获得执行的机会
	fmt.Println("main")
}


// -------------------当前goroutine退出------------------------
func goexit() {
	// goroutine A
	go func() {
		defer fmt.Println("defer")
		runtime.Goexit()
		fmt.Println("goroutine A")
	}()
	// goroutine B
	go func() {
		fmt.Println("goroutine B")
	}()
	time.Sleep(time.Second)
	fmt.Println("main")

}

// -----------------------runtime.GOMAXPROCS 指定使用多少个操作系统线程来执行当前的 go 语言程序-----------------
func goprocess() {

	runtime.GOMAXPROCS(1) // 只分配一个cpu, 永远都是一个线程从 0 打印到 9，然后再才打印另外一个线程。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("线程1：", i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("线程2：", i)
		}
	}()

	time.Sleep(time.Second)

}
