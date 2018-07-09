package main 
import "fmt"
func main(){
	var x,y,z int
	fmt.Println("Enter 3 numbers")
	fmt.Println("Enter number 1:")
	fmt.Scan(&x)
	fmt.Println("Enter number 2:")
	fmt.Scan(&y)
	fmt.Println("Enter number 3:")
	fmt.Scan(&z)
	if x<y && x<z{
		fmt.Println(x, "is smallest")
	}else if y<x && y<z{
		fmt.Println(y, "is smallest")
	}else{
		fmt.Println(z,"is smallest")
	}

}