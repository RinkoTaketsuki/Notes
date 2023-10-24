package example

import (
	"fmt"
	"math"
)

// 方法不过是把 this 当成一个参数。
// 分成值接收者和指针接收者两种情形。
// 方法和结构体的定义必须在同一 package 内。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Vertex 是 Abser，但 *Vertex 也是 Abser ?
type Abser interface {
	Abs() float64
}

// *Vertex 是 Scaler，但 Vertex 不是 Scaler
type Scaler interface {
	Scale(float64)
}

type Square struct {
	Edge float64
}

func (sq *Square) Scale(f float64) {
	sq.Edge *= f
}

func (sq Square) Area() float64 {
	return sq.Edge * sq.Edge
}

// 任何类型都属于 interface{}
func describe(i interface{}) {
	fmt.Printf("%#v\n", i)
}

func Run6() {
	fmt.Println("--------- Run 6 ---------")
	v := Vertex{1.0, 2.0}
	p := &v
	fmt.Printf("v.Abs(): %v\n", v.Abs())
	// 相当于 (&v).Scale(3.0)
	v.Scale(3.0)
	fmt.Printf("v: %v\n", v)
	// 相当于 (*p).Abs()
	fmt.Printf("p.Abs(): %v\n", p.Abs())
	p.Scale(2.0)
	fmt.Printf("p: %v\n", p)
	// 接口多态性，scaler 底层相当于存储了一个（类型，值）的二元组
	var scaler Scaler = &Vertex{114, 514}
	scaler.Scale(1.5)
	fmt.Printf("scaler: %T, %v\n", scaler, scaler)
	// 底层二元组的值可以是 nil
	var emptyVertex Vertex
	var nilVertex *Vertex
	var abser Abser = emptyVertex
	fmt.Printf("abser: %T, %v\n", abser, abser)
	abser = nilVertex
	fmt.Printf("abser: %T, %v\n", abser, abser)
	// 接口本身可以是 nil，此时不可调用接口中定义的方法，会触发空指针 panic
	var nilAbser Abser
	fmt.Printf("nilAbser: %T, %v\n", nilAbser, nilAbser)
	describe(&Vertex{-1, 2})
	var scaler2 Scaler = &Square{3}
	// 类型断言，若 scaler2 不是 *Square 则会触发 panic
	// 如果括号内的类型不是 Scaler 则无法通过编译
	square := scaler2.(*Square)
	fmt.Printf("square.Edge: %v\n", square.Edge)
	// 类型断言形式2，ok 表示断言成功，若断言失败会使 vertex 为 nil
	if vertex, ok := scaler2.(*Vertex); ok {
		fmt.Printf("vertex: %v\n", vertex)
	} else {
		fmt.Println("vertex is nil")
	}

}
