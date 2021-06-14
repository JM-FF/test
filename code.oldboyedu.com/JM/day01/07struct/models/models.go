package models

import "fmt"

type Person struct {
	Name string
	Age  int
	sal  float64
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

//Setage

func (p *Person) Setage(age int) {
	if age > 0 {
		p.Age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *Person) Getage() int {
	return p.Age
}

func (p *Person) Setsal(sal float64) {
	if sal > 0 && sal < 3000 {
		p.sal = sal
	} else {
		fmt.Println("薪资不在范围内！")
	}
}

func (p *Person) GetSal() float64 {
	return p.sal
}
