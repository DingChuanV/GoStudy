package main

import (
	"fmt"
)

/* 
	defer 关键字

		defer 关键字，用于在函数返回前执行一些操作
		例如，关闭文件、释放数据库连接等
		defer 关键字可以用来执行多个操作，每个操作之间用逗号隔开
		例如，下面的代码在函数返回前，先关闭文件，再释放数据库连接
		defer closeFile()
		defer releaseDBConnection()

		程序再运行时常常会创建一些临时资源，比如文件、数据库连接等
		这些资源在程序运行完成后，需要被释放
		defer 关键字，用来在函数返回前，释放这些临时资源

		清理共奏必须进行，不管函数有多少种推出方式。

*/

func main(){
	defer closeFile()
	defer releaseDBConnection()
}