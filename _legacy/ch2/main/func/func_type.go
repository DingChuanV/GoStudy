package main

import (
	"fmt"
	"strconv"
)

// 声明函数类型：接收两个 int，返回 int
type opFuncType func(int, int) int

// 四个简单算术实现
func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }
func div(a, b int) int {
	if b == 0 {
		// 在此示例中，除以 0 时返回 0（真实代码中应返回 error）
		return 0
	}
	return a / b
}

// compute 解析并计算一个三元表达式（如 ["1", "+", "2"]）
func compute(exp []string, ops map[string]opFuncType) {
	if len(exp) != 3 {
		fmt.Println("invalid expression:", exp)
		return
	}
	p1, err := strconv.Atoi(exp[0])
	if err != nil {
		fmt.Println("invalid operand:", exp[0])
		return
	}
	p2, err := strconv.Atoi(exp[2])
	if err != nil {
		fmt.Println("invalid operand:", exp[2])
		return
	}
	op := exp[1]
	opFunc, ok := ops[op]
	if !ok {
		fmt.Println("invalid operator:", op)
		return
	}
	res := opFunc(p1, p2)
	fmt.Printf("%d %s %d = %d\n", p1, op, p2, res)
}

func main() {
	// 使用自定义类型声明操作映射
	opMap := map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	// 示例用例（包含错误情况）
	examples := [][]string{
		{"10", "+", "5"},
		{"6", "/", "0"}, // 除以零（示例中返回 0）
		{"2", "^", "3"}, // 无效操作符
		{"a", "+", "1"}, // 无效操作数
	}

	for _, ex := range examples {
		compute(ex, opMap)
	}
}
