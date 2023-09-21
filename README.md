# Golang 个人笔记

好记性不如烂笔头。一些难以分类的笔记会单独写到 README 里面。

## 一个 import 关键字导入多个包

```go
import (
    "fmt"
    "time"
)
```

## 一次多个变量

```go
var (
    a int
    b string
)
var c, d, e int
var i, j int = 1, 2
l, m, n := 3, false, "ok"
```

## 参数列表的省略类型写法

```go
func add(x, y int) int {
    return x + y
}
```

## 返回值列表的省略名称写法

```go
func add(x, y int) int {
    return x + y
}
```

## 打印 Time 结构体

```go
fmt.println(time.Now()) // 2023-09-21 06:14:51.034412922 +0000 UTC m=+0.000050672
```

## 打印复数

```go
var z complex128 = cmplx.Sqrt(-5 + 12i)
fmt.Printf("Type: %T Value: %v\n", z, z) // Type: complex128 Value: (2+3i)
```

## 在函数体之外，任何语句都以关键字开头，故不能用 `:=` 声明变量

## Basic Types and Default Values（无需 make，值拷贝传参，没有 nil 的概念）

```go
bool // default: false

string // default: ""

// default: 0
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64 // default: 0

complex64 complex128 // default: (0+0i)
```

> 为了代码的跨平台性，最好不要用固定大小的 int 和 uint，这两个类型的大小通常总是机器的字长（32 位或 64 位）

## 字面值的默认类型

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

## 字面值、未定义类型的常量是可以被精确计算的（不限精度），但过高精度的数值存入变量后其值是不确定的

```go
package main

import "fmt"

const (
    // Create a huge number by shifting a 1 bit left 100 places.
    // In other words, the binary number that is 1 followed by 100 zeroes.
    Big = 1 << 100
    // Shift it right again 99 places, so we end up with 1<<1, or 2.
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))     // 21
    fmt.Println(needFloat(Small))   // 0.2
    fmt.Println(needFloat(Big))     // 1.2676506002282295e+29
}
```

## golang 所有实数类型之间的转换必须显式表示，除了整型提升
