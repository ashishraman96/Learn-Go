package main 
import (
"fmt"
"math"
)
func primeGenerator() func() int{
	i:=1
	isPrime:=true
	gen:=func() int{
		i=i+1
		j:=2
		tmp := math.Floor(math.Sqrt(float64(i)))
		//fmt.Println("i:", i, " sqrt:", int(tmp))
		for j<=int(tmp){
			if i%j==0{
				isPrime=false
				j++
			}
		}
		if isPrime{
			return i
		} 
		return -1
	}
	return gen
}

func main() {
	seq:=primeGenerator()
	for i:=0;i<100;i++{
		a:=seq()
		if a==-1{}else{
	fmt.Print(a, " ")
	}
}
}
