package main

import (
	"fmt"
	"strconv"
)

// var and const defination
var (
	a    string
	b, c bool
)
var d, e int
var f, g float32 = 0xf, 3.14

const untypedInt = 0xfffffffffffffffffffffffffff
const smallUntypedInt = untypedInt & 0xffff
const (
	untypedFloat   = -777777779999.777777779999
	untypedComplex = 777777779999.999999 + 77777777999.77777777999i
)

// type cast do not support integer truncate
var toInt = int32(smallUntypedInt)
var toUint16 = uint16(smallUntypedInt)

var toFloat32 = float32(untypedFloat)
var toFloat64 = float64(untypedFloat)

var toComplex64 = complex64(untypedComplex)
var toComplex128 = complex128(untypedComplex)

func printVars() {
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println(toInt, toUint16, toFloat32, toFloat64, toComplex64, toComplex128)
}

func add(x, y int) int {
	return x + y
}
func sub(x, y float64) (z float64) {
	z = x - y
	return
}

func addAndSub(x, y float64) (r1, r2 float64) {
	r2 = float64(add(int(x), int(y)))
	return sub(x, y), r2
}

func getFloatAndString(x int) (f float64, s string) {
	f = float64(x)
	s = strconv.Itoa(x)
	return
}
