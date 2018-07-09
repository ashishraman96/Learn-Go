package main 
import (
		"fmt"
		)
func main() {
	var rows, cols, i , j ,x int
	var a [][]int
	fmt.Print("Enter rows and columns:")
	fmt.Scan(&rows,&cols)
	//r:= make([]int,rows)
	//a = append(a,r)
	for i=0; i <rows ;i++ {
			
		for j=0; j < cols ; j++ {
			c:= make([]int,cols)
			fmt.Print("Enter array:")
			fmt.Scan(&x)
			c = append(c,x)
		}
		a = append(a,c)
	}
	fmt.Print(a)
	// i=0
	// j=0
	for i=0; i <rows ;i++{
		for j=0; j < cols ; j++{
			// fmt.Print(a[i][j]," ")
			// j=j+1
		} 
		fmt.Println(" ")
		i=i+1
	}
	

}