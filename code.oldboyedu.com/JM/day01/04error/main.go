package main

import (
	"fmt"
	"runtime"
	"time"
)

type ParseError struct {
	Filename string //文件名
	Line     int    //行号
}

//实现一个error接口
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)

}

func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func main() {

	start := time.Now()
	var e error
	fmt.Printf("%T", e)
	e = newParseError("main.go", 1)

	fmt.Println(e.Error())
	fmt.Println("---------------------------")

	ProtectRun(func() {
		fmt.Println("手动宕机前")
		panic(&panicContext{
			"手动触发",
		})

		fmt.Println("手动宕机后")

	})

	//指针造成错误访问

	ProtectRun(func() {
		fmt.Println("手动宕机前")

		var a *int
		*a = 1

		fmt.Println("手动宕机后")
	})

	// fmt.Println("----------exit-----------")
	// var b *int
	// *b = 1
	// defer exit()
	end := time.Since(start)
	fmt.Println(end)

}

type panicContext struct {
	function string
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		fmt.Println("处理函数")
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error", err)
		}
	}()
	entry()
}

func exit() {
	err := recover()
	fmt.Println("处理函数")
	switch err.(type) {
	case runtime.Error:
		fmt.Println("runtime error:", err)
	default:
		fmt.Println("error", err)
	}
}
