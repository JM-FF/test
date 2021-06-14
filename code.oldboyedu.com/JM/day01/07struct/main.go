package main

import (
	"fmt"

	"hello/models"
)

func main() {
	p := models.NewPerson("小明")

	p.Setage(20)
	p.Setsal(2500)
	fmt.Println(p)
	fmt.Println(p.GetSal())
	fmt.Println(p.Getage())
}
