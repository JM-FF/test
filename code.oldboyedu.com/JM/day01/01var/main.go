package main

import "fmt"

//声明变量
// var name string
// var age int
// var isOk bool

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "weideng"
	age = 20
	isOk = true
	var studentName string = "deng" //非全局变量的声明需要使用
	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name:%s\n", name)
	fmt.Println(age)
	fmt.Println(studentName)

	//类型推导
	var s1 = "20"
	fmt.Println(s1)

	//简短变量声明只能在函数中使用
	s2 := "xiaoxin"
	fmt.Println(s2)
	//同一个作用域不能声明一个变量

}
