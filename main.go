package main

import "fmt"

type Rectangle struct {
	Length, Width, Area, Perimeter int
}

func (a *Rectangle) RectangleArea() {
	a.Area = (a.Length * a.Width)
}

func (a *Rectangle) RectanglePerimeter() {
	a.Perimeter = (2 * (a.Length + a.Width))
}

func main() {
	TheRectangle := Rectangle{
		3,
		5,
		0,
		0,
	}

	fmt.Printf("Для прямоугольника со сторонами %v и %v \nПериметр = %v\nПлощадь = %v", TheRectangle.Length, TheRectangle.Width, TheRectangle.RectanglePerimeter(), TheRectangle.RectangleArea())
}
