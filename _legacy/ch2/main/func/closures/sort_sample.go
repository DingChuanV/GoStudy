package main

import (
	"fmt"
	"sort"
)

// 创建引用局部变量的闭包，并将其传递给另一个函数

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	people := []Person{
		{"Alice", "Smith", 30},
		{"Bob", "Doe", 25},
		{"Charlie", "Brown", 35},
	}
	fmt.Println(people)
	// 传递给 sort.Slice 函数的闭包 有两个参数 i和j 但在闭包 内部使用了切片 people ，因此就可以
	// 按LastName排序
	// 用计算机科学术语来说，people 被闭包捕获了

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
