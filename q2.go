package main

import ("fmt"; "math")

type Shape interface {
	area() float64
	perimeter() float64
}
type Rectangle struct {
	x1, y1, x2, y2 float64
}
type Circle struct {
	x, y, r float64
}
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2-x1
	b := y2-y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2*l + 2*w
}
func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}
func main() {

r := Rectangle{0, 0, 10,10}
c := Circle{0, 0, 5}
fmt.Println("perimeter of rect. is ", r.perimeter())
fmt.Println("perimeter of circle is ", c.perimeter())	
}
 
