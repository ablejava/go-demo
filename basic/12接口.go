package main

import "fmt"
// -------------定义接口-------------
type Phone interface {
	call()
}

// 定义iphone实现
type IPhone struct {

}
// 使用值接收参数
func (iphone IPhone) call() {
	fmt.Println("iphone call...")
}

// 定义小米phone实现
type XiaoMi struct {

}
func (xiaomi XiaoMi) call()  {
	fmt.Println("xiaomi call...")
}

func phoneCall() {
	var phone Phone			// 声明一个Phone类型的变量phone
	phone = new(IPhone)		// 实例化IPhone赋值给phone
	phone.call()			// 使用phone调用call方法，实际调用的是存储在phone变量里面的IPhone实例所实现的call方法
	phone = new(XiaoMi)		// 实例化XiaoMi赋值给phone
	phone.call()

}
// -----------------值接受者----------------------
func valueConsumer() {
	var phone Phone
	iphone1 := IPhone{}
	phone = iphone1		 			// 直接将值类型赋值给phone
	phone.call()
	iphone2 := &IPhone{}
	phone = iphone2					// 将指针类型&Phone赋值给phone, 会自动根据指针求值*iphone2，然后将其赋值给变量phone
	phone.call()

}

// 使用指针接受者实现接口
func (iphone *IPhone) pointCall() {
	fmt.Println("iphone call...")
}
// ---------------指针接收参数---------------------
func pointConsumer() {

	var phone Phone
	iphone1 := IPhone{}
	phone = iphone1		 			// 直接将值类型赋值给phone，报错
	phone.call()
	iphone2 := &IPhone{}
	phone = iphone2					// 将指针类型&Phone赋值给phone
	phone.call()

}

// -----------------------实现多个接口------------------------
type Game interface {
	play()
}
func (iphone IPhone) play() {
	fmt.Println("play game")
}