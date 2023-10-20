package main

import (
	"fmt"
	"strings"
)

func arrayAndSliceDefExamples() {
	// 数组的定义 1
	var a [2]string
	// 数组元素的赋值
	a[0] = "Hello"
	a[1] = "World"
	// 数组元素的引用
	fmt.Println(a[0], a[1])
	// 数组的定义 2
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// 切片的定义 1（左闭右开）
	var s1 []int = primes[1:4]
	fmt.Printf("%T, %v\n", s1, s1)
	// 切片是底层数组的引用
	s1[0] = 17
	fmt.Printf("primes: %T, %v\n s: %T, %v\n", primes, primes, s1, s1)
	// 切片的定义 2（相当于声明了一个隐含数组的 [:] 切片）
	s2 := []bool{true, false, true, true, false, true}
	fmt.Println(s2)
	// 切片的定义 3（花括号可嵌套）
	s3 := []struct {
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
	fmt.Println(s3)
	// 切片的定义 4（使用 make）
	ms1 := make([]int, 5)    // 相当于引用一个长度为 5 的数组，且 ms1 的 len 和 cap 均为 5
	ms2 := make([]int, 3, 5) // 相当于引用一个长度为 5 的数组，且 ms1 的 len 为 3，cap 为 5，引用底层数组的前 3 个元素
	fmt.Printf("ms1: %v\n", ms1)
	fmt.Printf("ms2: %v\n", ms2)
}

// 切片的长度就是它所包含的元素个数。
// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
func sliceLenAndCapExamples() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// 截取切片使其长度为 0
	s = s[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// 拓展其长度，注意底层数组始终只有一个，且不能扩展到超出底层数组的范围，否则会有 panic: runtime error
	s = s[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// 舍弃前两个值
	s = s[2:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliceZeroValueExamples() {
	var a [0]int
	fmt.Printf("a: %v\n", a)
	s := a[:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	var s2 []int
	if nil == s2 {
		fmt.Println("s2 is nil")
	}
}

func slice2dExamples() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
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
