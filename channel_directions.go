package main 
import "fmt"
func main() {
	c1:= make(chan int,1)
	c2:= make(chan int,1)

	go square(c1,c2)
	c1 <- 10
	result := <-c2
	fmt.Println("result : ", result)
}
func square(s,r chan int){
	d:= <-s
	r <-d*d
}