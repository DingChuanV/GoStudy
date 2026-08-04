package main

import "fmt"

/**
	匿名函数
	   不仅可以将函数赋值给 变量，还可以在函数中定义函数并将其复制给变量
	   这种函数称为匿名函数（anonymous function），因为它没有名字
	   匿名函数可以直接调用，也可以作为参数传递给其他函数
*/


var(
		add = func(i,j int) int {return i+j}
		sub = func(i,j int) int {return i-j}
		mul = func(i,j int) int {return i*j}
		div = func(i,j int) int {return i/j}
)

func main() {
	// 定义一个匿名函数并赋值给变量 f
	f := func(j int){
		fmt.Println(j)
	}
	for i := 0; i < 5; i++ {
		f(i) 
	}

	// 匿名函数 可以直接在定义后立即调用，无需赋值给变量。
	// 这种写法被称为 立即调用函数表达式（immediately invoked function expression，IIFE）
	for i := 0; i < 5; i++ {
		func(j int){
			fmt.Println(j)
		}(i)
	}

	// 通常开发不刻意采用这种写法，然在在一下场景中，匿名函数是非常有用的：
	// 1. defer 语句，延迟执行带有上下文状态的逻辑
	// 2. 启动 gorouti呢：通过轻量级并发执行代码块

	x := add(1, 2)
	fmt.Println(x)	

}

/**
	内部函数是没有名称的匿名函数： 在声明函数时，关键字func后面紧跟输入参数
	、返回值和左花括号。
*/

