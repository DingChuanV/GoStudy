package main

import "fmt"

func helloLen(s string) int {
	return len(s)
}

func sumRunes(s string) int {
	total := 0
	for _, r := range s {
		total += int(r)
	}
	return total
}

func main() {
	// handlers 映射命令名到具体的处理函数
	handlers := map[string]func(string) int{
		"len": helloLen,
		"sum": sumRunes,
	}

	inputs := []string{"hello", "世界"}
	for _, in := range inputs {
		fmt.Printf("input: %s\n", in)
		fmt.Printf("  len: %d\n", handlers["len"](in))
		fmt.Printf("  sum: %d\n\n", handlers["sum"](in))
	}
}
