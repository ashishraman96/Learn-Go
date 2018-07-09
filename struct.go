package main 
import "fmt"

type rectangle struct{
	length int
	breadth int
}

func (r1 rectangle) area() int {
	return r1.length*r1.breadth
}
 
func (r1 rectangle) isSquare() bool{
	if r1.length==r1.breadth{
		return true
	}else{
		return false
	}
}

func main(){
	r:=rectangle{2,3}
	a:=r.area()
	fmt.Println("Area=",a)
	f:= r.isSquare()
	if f==true{
		fmt.Println("Its a square")
	}else{
		fmt.Println("Its not a square")
	}
	r.scale(1,1)
	fmt.Print(r)

}
func (r1 *rectangle) scale(l,b int){
	r1.length+=l
	r1.breadth+=b
}