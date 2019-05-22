package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) //以指针为接收者的方法被调用时，接收者既能为值又能为指针，即便 v 是个值而非指针，带指针接收者的方法也能被直接调用
	ScaleFunc(&v, 10)
	fmt.Println(v.Abs())	//而以值为接收者的方法被调用时，接收者既能为值又能为指针
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))	//接受一个值作为参数的函数必须接受一个指定类型的值

	fmt.Println(v, p)
}
