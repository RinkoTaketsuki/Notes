package main

import (
	"fmt"
	"time"
)

func main() {
	h, i, j := (3 - 4i), "hello", true
	printVars()
	fmt.Printf("h: %v\n", h)
	fmt.Printf("i: %v\n", i)
	fmt.Printf("j: %v\n", j)
	fmt.Println(time.Now())
	fmt.Println(addAndSub(81, 49))
	fmt.Println(getFloatAndString(128))
	forLoopPattern1()
	forLoopPattern2()
	forLoopPattern3()
	fmt.Printf("pow(3, 2, 10): %v\n", pow(3, 2, 10))
	switchExamples()
	deferStackExamples()
	testVertex()
	arrayAndSliceDefExamples()
	sliceLenAndCapExamples()
	sliceZeroValueExamples()
}
