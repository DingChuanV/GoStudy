package main 

import (
	"fmt"
)

/**
	闭包 不仅可以用来再函数之间传递状态，还可以作为一个函数的返回值
	Go开发者是否频繁使用闭包，来创建函数的工厂函数，来创建不同的函数
	函数的工厂函数，返回值是一个函数，这个函数可以用来创建不同的函数
	例如，下面的代码创建了一个函数的工厂函数，用来创建不同的乘法函数
	每个乘法函数都有一个不同的基础值，但是都可以用来乘以不同的因子
	 还通过 defer 关键字，利用闭包来实现资源清理

	 高阶函数，简单来说 高阶函数是指以函数作为输入参数或返回值的函数
*/

func makeMult(base int) func(int) int{
	return func(factor int) int{
		return base * factor
	}
}

func main(){
	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i :=0;i<3;i++{
		fmt.Println(twoBase(i),threeBase(i))
	}
}