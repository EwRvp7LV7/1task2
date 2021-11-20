package module

import (
	"fmt"
	"math"
	//"strings"
)

type Circle struct {
	x, y, r float64
}

// //function
// func circleArea(c *Circle) float64 { //pointer, not copy value!
// 	return math.Pi * c.r * c.r
// }

//method
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

//Объединяем объекты с одинаковыми методами как Shape
//Shape это псевдообъект, объединение объектов с одинаковыми методами
//но в реализации вместо него ставятся реальные объекты
type Shape interface {
	area() float64
	//area1() float64 //чтобы принадлежать интерфейсу тип должен иметь все методы итф плюс может другие свои
}

//компилятор сам определяет какая area принадлежит какому объекту
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func CalcShapes() {
	c := Circle{0, 0, 5}
	// fmt.Println(circleArea(&c))
	// fmt.Println(c.area())
	r := Rectangle{0, 0, 10, 10}
	// fmt.Println(r.area())
	fmt.Println(totalArea(&c, &r)) //здесь если r не имеет area() компилятор не пустит

}
