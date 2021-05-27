package main

import "fmt"

//常量
//定义以后不能修改
const pi = 3.1415926

//批量声明
const (
	statusOk = 200
	notFound = 404
)

//某一行没有赋值默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

//iota

const (
	a1 = iota
	a2
	a3
)

const (
	b1 = iota
	b2
	_
	b3
)

type Circle struct {
	radius float64
}

type Student struct {
	num   int
	name  string
	class string
}

func main() {
	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
	fmt.Println("a3: ", a3)

	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)
	fmt.Println("b3: ", b3)
	var sum int = 0

	//i变量旨在for中生效
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	if age := 20; age >= 18 {
		fmt.Println("已经成年")
	} else {
		fmt.Println("未成年")
	}

	str := []string{"小明", "小张", "小王"}

	for _, s1 := range str {
		fmt.Println(s1)
	}

	var i int
	var j int
	for i = 1; i < 10; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d   ", i, j, i*j)
		}
		fmt.Println()
	}

	var num1 int = 10
	var num2 int = 5
	fmt.Println(max1(num1, num2))
	fmt.Println("---------------交换----------")
	fmt.Println(num1)
	fmt.Println(num2)
	swap1(num1, num2)
	fmt.Println(num1)
	fmt.Println(num2)
	swap2(&num1, &num2)
	fmt.Println(num1)
	fmt.Println(num2)

	fmt.Println("------------函数实参-------------")
	getmax := func(n1, n2 int) int {
		return max1(n1, n2)
	}

	println(getmax(num1, num2))

	fmt.Println("-----------------闭包-------------")
	nextNumber := getnum()
	fmt.Println("nextnumber1: ")
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	nextNumber2 := getnum()
	fmt.Println("nextnumber2: ")
	fmt.Println(nextNumber2())
	fmt.Println(nextNumber2())
	fmt.Println("----------------方法-----------------")
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积： ", c1.getArea())

	fmt.Println("---------------数组-----------------")
	var a = []int{10, 20, 30}
	//var b []int
	var c []*int
	for i := 0; i < 3; i++ {
		c = append(c, &a[i])
	}

	fmt.Println(*&a)
	fmt.Println(c)
	fmt.Println(*&c)
	fmt.Println("-------------------")
	for i := 0; i < len(c); i++ {
		fmt.Println(*c[i])
	}
	fmt.Println("--------------结构体------------")
	var student1 Student
	var student2 = Student{num: 2, name: "xiaoming", class: "二班"}
	student1 = Student{num: 1, name: "xioahong", class: "一班"}
	student3 := Student{num: 3, name: "xiaohei", class: "三班"}
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(student3)
	student1.name = "张三"
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(student3)
	fmt.Println("-----------结构体指针----------------")
	var studentPointer *Student
	studentPointer = &student1
	fmt.Println(*studentPointer)
	fmt.Println((*studentPointer).name)
	fmt.Println("------------------------------")
	fmt.Println(*&studentPointer.num)
	fmt.Println(*&studentPointer.name)
	fmt.Println(*&studentPointer.class)
	fmt.Println("----------------------------------")
	fmt.Println(studentPointer.num)
	fmt.Println(studentPointer.name)
	fmt.Println(studentPointer.class)
	fmt.Printf("%T\n", studentPointer)
	fmt.Println("-------------------------")
	fmt.Println(student1.num)
	fmt.Println(student1.name)
	fmt.Println(student1.class)

	fmt.Println("---------------------切片---------------")
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2)
	printSlice(numbers)

	numbers = append(numbers, 2)
	printSlice(numbers)

	numbers = append(numbers, 3, 4, 5)
	printSlice(numbers)

	numbers = append(numbers, 2)
	printSlice(numbers)

	numbers = append(numbers, 2)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
	fmt.Println("--------切片的扩容方式是二倍的增长方式---------")

	var m1 [3]*int
	var m2 = []int{0, 1, 2, 3, 4}
	fmt.Println(m2[0:2])
	m2 = append(m2, 7)
	fmt.Println(m1)

	m1[1] = &m2[1]
	fmt.Println(*m1[1])
	//var m1 []*int
	//m1 = append(m1,&m2[1])
	fmt.Println("--当数组指定大小以后就不能当作切片使用也不能使用append函数---指针数组同理--")

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func (c Circle) getArea() float64 {
	return pi * c.radius * c.radius
}

func max1(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func swap1(num1 int, num2 int) {
	var tmp int
	tmp = num1
	num1 = num2
	num2 = tmp
}

func swap2(num1 *int, num2 *int) {
	var tmp int
	tmp = *num1
	*num1 = *num2
	*num2 = tmp
}

//闭包
func getnum() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
