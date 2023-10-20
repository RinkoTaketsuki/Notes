package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func forLoopPattern1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forLoopPattern2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func forLoopPattern3() {
	sum := 1
	for {
		if sum > 10 {
			break
		}
		sum += sum
	}
	fmt.Println(sum)
}

// sentence before if condition
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	const EPSILON = 0.000001
	z := 1.0
	delta := math.MaxFloat64
	for math.Abs(delta) > EPSILON {
		delta = (z*z - x) / (2 * z)
		z -= delta
	}
	return z
}

func getWindows() string {
	fmt.Println("f is called")
	return "Windows"
}

func switchExamples() {
	fmt.Println("Go runs on")
	linux := "linux"
	// "switch" statement does not need break at each case,
	// if you need the C-like "switch", use "fallthrough"
	switch os := runtime.GOOS; os {
	case "darwin": // can be non-integer
		fmt.Println("OS X.")
	case linux: // can be non-const
		fmt.Println("Linux.")
	case getWindows(): // if os matches the conditions above, getWindows() can not be called
		fmt.Println("Windows")
	default:
		fmt.Printf("%s.\n", os)
	}

	// switch true pattern
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func itoaWrapper(i int) int {
	fmt.Printf("itoaWrapper is called, i is %v\n", i)
	return i
}

func deferStackExamples() {
	fmt.Println("defer begin")
	for i := 0; i < 3; i++ {
		defer fmt.Println(itoaWrapper(i))
	}
	fmt.Println("defer end")
}
