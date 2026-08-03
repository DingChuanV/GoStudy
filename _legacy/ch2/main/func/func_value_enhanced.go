package main

import (
	"fmt"
	"strconv"
)

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func addOp(a, b int) (int, error) { return a + b, nil }
func subOp(a, b int) (int, error) { return a - b, nil }
func mulOp(a, b int) (int, error) { return a * b, nil }
func divOp(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// computeExpression 解析并计算一个三元表达式
func computeExpression(exp []string, opMap map[string]func(int, int) (int, error)) {
	if len(exp) != 3 {
		fmt.Println("invalid expression:", exp)
		return
	}

	p1, err := strconv.Atoi(exp[0])
	if err != nil {
		fmt.Println("invalid operand:", exp[0])
		return
	}

	op := exp[1]
	var binOp func(int, int) (int, error)
	var ok bool
	if binOp, ok = opMap[op]; !ok {
		fmt.Println("invalid operator:", op)
		return
	}

	p2, err := strconv.Atoi(exp[2])
	if err != nil {
		fmt.Println("invalid operand:", exp[2])
		return
	}

	res, err := binOp(p1, p2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%d %s %d = %d\n", p1, op, p2, res)
}

func main() {
	// 演示函数作为值（保留原示例）
	var myFunc func(string) int
	myFunc = f1
	fmt.Println(myFunc("hello"))

	myFunc = f2
	fmt.Println(myFunc("hello"))

	// 改进后的运算映射，支持错误返回
	opMap := map[string]func(int, int) (int, error){
		"+": addOp,
		"-": subOp,
		"*": mulOp,
		"/": divOp,
	}

	// 包含一些正常与错误用例，便于学习
	expressions := [][]string{
		{"1", "+", "2"}, // 正常
		{"3", "-", "4"}, // 正常
		{"5", "*", "6"}, // 正常
		{"7", "/", "0"}, // 除以零错误
		{"x", "+", "1"}, // 无效操作数
		{"9", "?", "1"}, // 无效操作符
	}

	for _, exp := range expressions {
		computeExpression(exp, opMap)
	}
}
