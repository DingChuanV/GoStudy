package main

import (
	"fmt"
)

/*
	嵌入不是继承
		- Go 的嵌入不是继承。
		- Outer 没有“重写” Inner 的方法。
		- Outer.IntPrinter 只是和 Inner.IntPrinter 同名，并不会影响 Inner.Double() 内部的调用。
*/

type Inner struct{
	A int
}

func (i Inner) IntPrinter(val int) string{
	return fmt.Sprintf("Inner %d", val)
}

func (i Inner) Double() string{
	return i.IntPrinter(i.A * 2)
}

type Outer struct{
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string{
	return fmt.Sprintf("Outer %d", val)
}


func main(){
	o := Outer{
		Inner: Inner{
			A : 10,
		},
		S : "Hello",
	}
	fmt.Println(o.Double())
	fmt.Println(o.IntPrinter(10))
	fmt.Println(o.Inner.IntPrinter(30))
}