package main

import (
	"fmt"
	"math"
)

// 1、定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格

// Shape 接口定义了形状的基本行为
// Area() 计算面积
// Perimeter() 计算周长
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 表示矩形
// Width: 宽度
// Height: 高度
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 计算矩形的面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形的周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 表示圆形
// Radius: 半径
type Circle struct {
	Radius float64
}

// Area 计算圆形的面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆形的周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// PrintShapeInfo 打印形状的信息
func PrintShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f\n", s.Area())
	fmt.Printf("周长: %.2f\n", s.Perimeter())
}

// 2、使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

// Person 表示一个人的基本信息
type Person struct {
	Name string
	Age  int
}

// Employee 表示员工信息，通过组合 Person 结构体
type Employee struct {
	Person      // 嵌入 Person 结构体，实现组合
	EmployeeID  string
}

// PrintInfo 打印员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n")
	fmt.Printf("姓名: %s\n", e.Name)        // 可以直接访问 Person 的字段
	fmt.Printf("年龄: %d\n", e.Age)         // 可以直接访问 Person 的字段
	fmt.Printf("工号: %s\n", e.EmployeeID)
}

func main() {
	// 测试形状接口
	fmt.Println("=== 形状接口测试 ===")
	rect := Rectangle{
		Width:  5,
		Height: 3,
	}
	fmt.Println("矩形信息:")
	PrintShapeInfo(rect)

	circle := Circle{
		Radius: 4,
	}
	fmt.Println("\n圆形信息:")
	PrintShapeInfo(circle)

	// 测试结构体组合
	fmt.Println("\n=== 结构体组合测试 ===")
	employee := Employee{
		Person: Person{
			Name: "张三",
			Age:  28,
		},
		EmployeeID: "EMP001",
	}
	employee.PrintInfo()
}
