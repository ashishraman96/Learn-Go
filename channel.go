package main 
import (
"fmt"
"time"
)


// func main(){
// 	a:=2
// 	b:=3
// 	c:= make(chan int)
// 	go add(a,b,c)
// 	r:= <-c
// 	fmt.Println("sum:",r)
// }

// func add(a,b int, c chan int){
// 	s:=a+b
// 	time.Sleep(time.Second)
// 	c <- s
// }
func main(){
	a:= []int {1,2,3,4,5,6,7,8,9,10}
	n:= len(a)
	c1:= make(chan int)
	c2:= make(chan int)
	go addn(a[:n/2],c1)
	go addn(a[n/2:],c2)
	select{ //to select whichever finishes first 
	case r1:= <-c1: 
		fmt.Println("r1:",r1)
	case r2:= <-c2:
		fmt.Println("r2:",r2)
	case t:= <- time.After(time.Second*10):
		fmt.Println(t)
	}
}
func addn(s []int, c chan int){
	r:=0
	for _,x := range s{
		time.Sleep(time.Second)
		r=r+x
	}
	c <- r
}