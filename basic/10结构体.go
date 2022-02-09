package main

import "fmt"

// -----------------定义结构体----------------------
//结构体与Java中的类很像
type Student struct {
	Name string

	Age int

	Gender string

	Grade string
}
// ------------------初始化结构体----------------------
func init() {
	// 1 这种方式给成员变量赋值需要全部赋值，不能只给其中部分赋值，并且有序，
	stu1 := Student{"张三", 18, "男", "大一"}
	// 2 这种方式可以部分赋值
	stu2 := Student{
		Name: "李四",
		Age: 20,
		Gender: "女",
		Grade: "大三",
	}
	// 3
	stu3 := Student{}
	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu3)

	// 打印结构体的值
	stu := Student{"张三", 18, "男", "大一"}
	fmt.Println("学生姓名：", stu.Name)
	fmt.Println("学生年龄：", stu.Age)
	fmt.Println("学生性别：", stu.Gender)
	fmt.Println("学生年级：", stu.Grade)
	// 修改结构体中的值
	stu.Name = "李四"
	stu.Age = 19
	stu.Gender = "女"
	stu.Grade = "大二"
	// 打印修改后的结构体的值
	fmt.Println("-----------------")
	fmt.Println("学生姓名：", stu.Name)
	fmt.Println("学生年龄：", stu.Age)
	fmt.Println("学生性别：", stu.Gender)
	fmt.Println("学生年级：", stu.Grade)


}

// ---------------定义方法--------------------
func (stu Student) print()  {
	fmt.Println("姓名", stu.Name)
	fmt.Println("年龄", stu.Age)
	fmt.Println("性别", stu.Gender)
	fmt.Println("年级", stu.Grade)
}

// --------------------定义结构体指针---------------------------
func structPoint() {
	// 声明结构体并赋值
	stu := Student{"张三", 18, "男", "大一"}
	// 定义结构体指针
	var stuPointer *Student
	// 将结构体地址赋值给指针
	stuPointer = &stu
	// 结构体指针通过点（.）操作符访问成员变量
	fmt.Println("姓名：", stuPointer.Name)
	fmt.Println("年龄：", stuPointer.Age)
	fmt.Println("性别：", stuPointer.Gender)
	fmt.Println("年级：", stuPointer.Grade)
}