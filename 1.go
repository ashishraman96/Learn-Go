package main 
import "fmt"
func main(){
	var n int
	fmt.Println("Enter a number:")
	fmt.Scan(&n)
	if n%2==0{
		fmt.Printf("%d is even",n)
	}else{
		fmt.Println(n," is odd \n")
	}
}