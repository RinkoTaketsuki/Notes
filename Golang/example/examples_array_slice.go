package example

import (
	"fmt"
	"strings"
)

func arrayAndSlice() {
	// 数组的定义 1
	var a [2]string
	// 数组元素的赋值
	a[1] = "World"
	// 数组元素的引用
	fmt.Printf("a[0]: %v\n", a[0])
	fmt.Printf("a[1]: %v\n", a[1])

	// 数组的定义 2
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 切片的定义 1（左闭右开）
	var primes1to4 []int = primes[1:4]
	fmt.Printf("primes1to4: %T, %v\n", primes1to4, primes1to4)

	// 切片是底层数组的引用
	primes1to4[0] = 17
	fmt.Printf("primes: %T, %v\n", primes, primes)

	// 切片的定义 2（相当于声明了一个隐含数组的 [:] 切片）
	boolValues := []bool{true, false, true, true, false, true}
	fmt.Printf("boolValues: %v\n", boolValues)

	// 切片的定义 3（花括号可嵌套）
	structValues := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Printf("structValues: %v\n", structValues)

	// 切片的定义 4（使用 make）
	ms1 := make([]int, 5)    // 相当于引用一个长度为 5 的数组，且 ms1 的 len 和 cap 均为 5
	ms2 := make([]int, 3, 5) // 相当于引用一个长度为 5 的数组，且 ms1 的 len 为 3，cap 为 5，引用底层数组的前 3 个元素
	fmt.Printf("make([]int, 5): %v\n", ms1)
	fmt.Printf("make([]int, 3, 5): %v\n", ms2)
}

func printLenAndCap(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// 切片的长度就是它所包含的元素个数。
// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。(注意不是底层数组的长度)
func sliceLenAndCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printLenAndCap(s)
	// 截取切片使其长度为 0
	s = s[:0]
	printLenAndCap(s)
	// 拓展其长度，注意底层数组始终只有一个，且不能扩展到超出底层数组的范围，否则会有 panic: runtime error
	s = s[:4]
	printLenAndCap(s)
	// 舍弃前两个值
	s = s[2:]
	printLenAndCap(s)
}

// 可以声明长度为 0 的数组，以及长度为 0 的切片。
// 注意区分 nil 和长度为 0 的切片。
func sliceZeroValue() {
	var a [0]int
	fmt.Printf("var a [0]int, a: %T, %v\n", a, a)
	s := a[:]
	fmt.Printf("a[:]: len=%d cap=%d %v\n", len(s), cap(s), s)
	var s2 []int
	if nil == s2 {
		fmt.Println("var s2 []int, s2 is nil")
	}
}

// 切片的元素也可以是切片
func slice2d() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceAppend() {
	var s []int
	printLenAndCap(s)

	// 添加一个空切片
	s = append(s, 0)
	printLenAndCap(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printLenAndCap(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printLenAndCap(s)
}

func arrayAndSliceForRange() {
	var pow1 = [8]int{1, 2, 4, 8, 16, 32, 64, 128}
	// 通常形式的 for range
	for i, v := range pow1 {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	pow2 := make([]int, 10)
	// 可以省略第二个值，也可以写成 i, _
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	// 省略第一个值需要显式地使用 _
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}

// 二维切片的分配、引用和 for range
func iota2d(size int) (s [][]int) {
	s = make([][]int, size)
	prev := 0
	for i := range s {
		s[i] = make([]int, size)
		s[i][0] = prev
		prev++
		for j := range s[i][:len(s[i])-1] {
			s[i][j+1] = s[i][j] + 1
		}
	}
	return
}

func Run4() {
	fmt.Println("--------- Run 4 ---------")
	arrayAndSlice()
	sliceLenAndCap()
	sliceZeroValue()
	slice2d()
	sliceAppend()
	arrayAndSliceForRange()
	fmt.Printf("iota2d(5): %v\n", iota2d(5))
}
