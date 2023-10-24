package example

import "fmt"

type Coordinate struct {
	X, Y int
}

var (
	v1 = Coordinate{1, 2}  // 创建一个 Coordinate 类型的结构体
	v2 = Coordinate{X: 1}  // Y:0 被隐式地赋予
	v3 = Coordinate{}      // X:0 Y:0
	p  = &Coordinate{1, 2} // 创建一个 *Coordinate 类型的结构体（指针）
)

func Run3() {
	fmt.Println("--------- Run 3 ---------")
	fmt.Println(v1, p, v2, v3, v1.X, p.Y) // p.Y 等价于 (*p).Y
}
