package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	// 检查命令行参数数量，如果小于2则说明没有指定要读取的文件
	if len(os.Args) < 2{
		log.Fatal("no file specified")
	}
	// 打开命令行参数中指定的第一个文件
	f ,err := os.Open(os.Args[1])

	// 如果打开文件出错，直接终止程序并打印错误信息
	if err != nil{
		log.Fatal(err)
	}
	// 延迟关闭文件，确保程序退出前资源被释放
	defer f.Close()

	// 创建2048字节的缓冲区，用于存储每次读取的文件内容
	data := make([]byte,2048)
	// 循环读取文件内容，直到读取完毕或出错
	for{
		// 从文件中读取最多2048字节的数据，返回读取到的字节数和可能的错误
		count,err := f.Read(data)
		// 将读取到的字节内容写入标准输出（控制台）
		os.Stdout.Write(data[:count])
		// 处理读取过程中出现的错误
		if err != nil{
			// 如果是到达文件末尾的错误，正常退出循环
			if err == io.EOF{
				break
			}
			// 其他错误则终止程序并打印错误信息
			log.Fatal(err)
		}
	}
}