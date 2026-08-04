package main

import "fmt"

func main() {

	/* 
		编写两个函数
			UpdateSlice 接收一个 []string 切片和一个 string 作为参数，
			将传入的 []string 切片的最后一个位置设置为传入的 string

			GrowSlice 接收一个 []string 切片和一个 string 作为参数，
			将 string 追加到切片中（使用 append）
	*/
	
	// ====== 测试 UpdateSlice ======
	fmt.Println("===== UpdateSlice ======")
	s1 := []string{"apple", "banana", "cherry"}
	fmt.Printf("修改前: %v (len=%d, cap=%d)\n", s1, len(s1), cap(s1))

	UpdateSlice(s1, "date")
	fmt.Printf("修改后: %v (len=%d, cap=%d)\n", s1, len(s1), cap(s1))

	/*
	 * UpdateSlice 的行为：
	 * 切片是引用类型（包含指向底层数组的指针、长度、容量）。
	 * UpdateSlice 直接通过索引修改了底层数组中的元素，
	 * 所以 main 中可以看到变化 —— 因为 s1 和函数内的 slice
	 * 共享同一个底层数组。
	 */

	// ====== 测试 GrowSlice ======
	fmt.Println("\n===== GrowSlice ======")
	s2 := []string{"hello", "world"}
	fmt.Printf("追加前: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	s2 = GrowSlice(s2, "go")
	fmt.Printf("追加后: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	// 继续追加，观察扩容行为
	s2 = GrowSlice(s2, "java")
	fmt.Printf("再追加: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	s2 = GrowSlice(s2, "rust")
	fmt.Printf("再追加: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	s2 = GrowSlice(s2, "python")
	fmt.Printf("再追加: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	/*
	 * 为什么有些变化不可见？
	 *
	 * 对于 UpdateSlice：可见。因为修改的是已有元素（通过索引直接写底层数组），
	 * 不改变切片头部的 len/cap/ptr，所以调用方的切片变量仍然指向同一个底层数组。
	 *
	 * 对于 GrowSlice：需要返回新的切片！因为 append 可能触发扩容：
	 *   - 如果 cap 足够：append 在原数组上追加，main 中的切片能看到变化
	 *     （但前提是 main 用返回值更新了自己的切片变量）
	 *   - 如果 cap 不够：append 分配新数组、复制旧数据、追加新元素，
	 *     返回一个指向新数组的切片。如果不接收返回值，main 中的旧切片
	 *     仍然指向旧数组，就看不到新追加的元素。
	 *
	 * 关键结论：append 可能返回一个新的切片头部，**必须接收其返回值**。
	 */
}

func UpdateSlice(s []string, str string) {
	if len(s) > 0 {
		s[len(s)-1] = str
	}
}

// 
func GrowSlice(s []string, str string) []string {
	return append(s, str)
}