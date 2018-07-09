package main 
import (
	"fmt"
	)

type rectangle struct{
	Length float32
	Breadth float32
}
func (r *rectangle) area() float32 {
	return r.Length * r.Breadth
}

type triangle struct{
	Base float32
	Height float32
}

func (t *triangle)area() float32{
	return 0.5 * t.Base * t.Height
} 

type shape interface{
	area() float32
}

func GetArea(Shape shape) float32{
	return Shape.area()
}


func main(){
	r := &rectangle{Length:2, Breadth:3}
	t:=  &triangle{Base:2, Height:4}
	ra:=r.area()
	ta:=t.area()
	fmt.Println("rectangle Area ra:", ra, " \t triangle area ta :",ta)
	rai:= GetArea(r)
	tai:= GetArea(t)
	fmt.Println("rectangle Area ra interface:", rai, " \t triangle area ta interface:",tai)
}