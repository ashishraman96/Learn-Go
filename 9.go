package main 
import (
	"fmt" 
	)

func main() {
	var r,c int
	var a,b,z [][]int
	fmt.Print("Enter rows and columns fot M1:")
	fmt.Scan(&r,&c)
	readMatrix(&a,r,c)
	fmt.Print("Enter rows and columns fot M2:")
	fmt.Scan(&r,&c)
	readMatrix(&b,r,c)
	add(&a,&b,&z,r,c)
	display(&z,r,c)
}

func readMatrix(p *[][]int,r,c int){	
	var x int
	fmt.Print("Enter values for matrix:")
	for i:=0;i<r;i++{
		var col []int
		for j:=0;j<c;j++{
				fmt.Scan(&x)
				col = append(col,x)
		}
		*p = append(*p,col)
	}
}

func add(a,b,z *[][]int, r, c int){
	for i:=0;i<r;i++{
		var col []int
		for j:=0;j<c;j++{
			col=append(col,(*a)[i][j]+(*b)[i][j])
		}
		*z = append(*z,col)
	}
}

func display(a *[][]int, r, c int){
	for i:=0;i<r;i++{
		for j:=0;j<c;j++{
			fmt.Print((*a)[i][j])
		}
		fmt.Println(" ")
	}
}