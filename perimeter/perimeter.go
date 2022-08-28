package perimeter

type Rectangle struct {
	Width  float64
	Height float64
}

//不同的包可以有函数名相同的函数，可以为多个类型定义各自的方法
//方法和函数的区别：
//方法是特定类型的实例调用的，但函数可以随时被调用
func (r Rectangle) Area() float64 {
	return 0
}

func (c Circle) Area() float64 {
	return 0
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(circle Circle) float64 {
	return circle.Radius
}
