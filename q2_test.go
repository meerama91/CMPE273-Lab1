package geo

import "testing"

func testPerimeter(t *testing.T) {
r := Rectangle{0, 0, 10,10}
per := r.perimeter()

if per != 40 {
    t.Error("Wrong perimeter of rectangle")
}
}

//func testPerimeterCircle(t *testing.T) {
//c := Circle{0, 0, 5}
//perc := c.perimeter()

//if perc != 21 {
//    t.Error("21 not received as perimeter of circle!")
//}
//}
