package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

/* 
	1000 万 Person 与 GC 观察
		1.创建 1000 万个 Person 切片会分配大量堆内存（约 1GB+）
	可通过调整 GOGC 环境变量观察 GC 对执行时间的影响：
		GOGC=off → 最快但内存飙升
		GOGC=200 → 较少 GC、较高内存
		GOGC=50 → 频繁 GC、低内存
GODEBUG=gctrace=1 可查看每次 GC 的详细 trace
代码中已内置内存统计输出（Alloc / TotalAlloc / Sys / NumGC）
*/

func main() {
	// 打印初始 GOGC 设置
	gogc := os.Getenv("GOGC")
	if gogc == "" {
		gogc = "100"
	}
	fmt.Printf("当前 GOGC 值: %s\n", gogc)

	// 记录开始时间
	start := time.Now()
	
	// 创建包含1000万个 Person 的切片
	persons := make([]Person, 10_000_000)

	// 填充切片
	for i:= range persons {
		persons[i] = Person{
			FirstName: "名字",
			LastName:  "姓氏",
			Age:       i % 100,
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("创建 1000 万个 Person 耗时: %s\n", elapsed)
	fmt.Printf("切片大小: len=%d, cap=%d\n", len(persons), cap(persons))

	// 手动触发 GC 并打印统计信息
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("\n--- 内存统计 ---\n")
	fmt.Printf("堆内存分配 (Alloc): %.2f MB\n", float64(memStats.Alloc)/1024/1024)
	fmt.Printf("总堆内存分配 (TotalAlloc): %.2f MB\n", float64(memStats.TotalAlloc)/1024/1024)
	fmt.Printf("系统内存申请 (Sys): %.2f MB\n", float64(memStats.Sys)/1024/1024)
	fmt.Printf("GC 次数 (NumGC): %d\n", memStats.NumGC)

	// 打印 GOGC 说明
	fmt.Printf("\n=== GOGC 与 GC 行为说明 ===\n")
	fmt.Printf(`
GOGC 环境变量控制 Go 垃圾回收器的触发频率：
- GOGC=off    : 禁用 GC（仅用于调试，生产环境不要用）
- GOGC=100    : 默认值，当堆内存增长 100% 时触发 GC
- GOGC=200    : 堆增长 200% 才触发 GC（更少的 GC 次数，但更高内存占用）
- GOGC=50     : 堆增长 50% 就触发 GC（更频繁的 GC，更低内存占用）

GODEBUG=gctrace=1 可以在运行时查看每次 GC 的详细信息。

实验建议：
  1. go run main.go                          # 使用默认 GOGC=100
  2. GOGC=off go run main.go                 # 禁用 GC，观察内存持续增长
  3. GOGC=200 go run main.go                 # 降低 GC 频率
  4. GODEBUG=gctrace=1 go run main.go        # 查看 GC trace

如果以 10,000,000 的容量创建 []Person 切片：
- 每个 Person 约 48 字节（两个 string 头部各 16 字段 + Age 8 字节 + 对齐）
- 但 string 内容会额外分配在堆上
- 10M 个 Person 的原始数据约 480MB+，加上字符串内容可能超过 1GB
- 这会导致显著的 GC 压力，可以通过调整 GOGC 观察对执行时间和内存的影响
`)

	// 可选：根据命令行参数设置不同的 GOGC 值进行对比测试
	if len(os.Args) > 1 {
		fmt.Printf("\n检测到命令行参数，可尝试不同 GOGC 值:\n")
		fmt.Printf("  GOGC=100 go run main.go\n")
		fmt.Printf("  GOGC=off  go run main.go\n")
		fmt.Printf("  GODEBUG=gctrace=1 GOGC=200 go run main.go\n")
	}
}
