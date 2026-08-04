package main

import (
	"fmt"
)

type Person struct{
	FirstName string
	LastName string
	Age int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName: lastName,
		Age: age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName: lastName,
		Age: age,
	}
}

func main() {
	// Person 结构体与逃逸分析

	// 调用 MakePerson，返回值类型
	p1 := MakePerson("张", "三", 25)
	fmt.Printf("MakePerson 返回值: %+v\n", p1)

	// 调用 MakePersonPointer，返回指针类型
	p2 := MakePersonPointer("李", "四", 30)
	fmt.Printf("MakePersonPointer 返回值: %+v\n", *p2)
	fmt.Printf("MakePersonPointer 指针地址: %p\n", p2)

	/*
	 * 关于逃逸到堆上的问题：
	 *
	 * MakePerson 返回的是 Person 值类型。如果调用方只是使用这个值（比如赋给局部变量），
	 * Person 通常分配在栈上，不会逃逸。
	 *
	 * MakePersonPointer 返回的是 *Person 指针。由于指针的生命周期超出了函数作用域
	 * （它被返回给了调用者），Go 编译器**必须**将这个 Person 分配在堆上，
	 * 否则函数返回后栈帧被销毁，指针就会指向无效内存。
	 *
	 * 可以通过 `go build -gcflags="-m"` 验证：
	 *   - MakePerson 中的 Person 通常不会逃逸（moved to heap: false）
	 *   - MakePersonPointer 中的 Person 会逃逸到堆上（moved to heap: true）
	 *
	 * 这就是"逃逸分析"的核心：编译器决定变量应该分配在栈还是堆上。
	 */

	
}