package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	/*
		谨慎使用指针
		1. 指针的使用会增加程序的复杂度，容易引入bug
		2. Go语言中，函数参数传递是值传递，传递指针可以避免拷贝大对象，提高性能
		3. 但是，过度使用指针会导致代码难以理解和维护
		4. 在Go语言中，推荐尽量使用值传递，只有在必要时才使用指针
	*/


	// 通常应该只在函数期望一个接口(interface)时,使用指针参数来修改变量，这种变成在处理json数据尤为常见
	f := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	err := json.Unmarshal([]byte(`{"name": "Bob", "age": 25}`), &f)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("Unmarshalled data: %+v\n", f)

	/* 
		Unmarshal 函数从一个包含 JSON 的字节切片中一个变量，该函数被声明为接受一个字节切片和一个any类型的参数。
		传递给any的参数的值必须是一个指针,否则，将会返回一个错误。

		json.Unmarshal 函数为什么要求传入指针，而不是直接返回一个结构体值？
			- 首先 函数 Unmarshal 的设计是为了能够修改传入的变量的值，而不是返回一个新的值。
			- 其次，传入指针可以避免在函数内部创建一个新的结构体实例，从而节省内存和提高性能。
			- 最后，传入指针可以让函数直接修改原始变量的值，而不是返回一个新的值，这样可以避免不必要的拷贝操作。
	*/

}

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo() (Foo, error) {
	// 1. 创建一个Foo对象
	foo := Foo{
		Field1: "value1",
		Field2: 42,
	}
	return foo, nil
}

