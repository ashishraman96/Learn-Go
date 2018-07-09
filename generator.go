package main 
import "fmt"
func sequence(seed int) func() int{
	var j int
	i:=0
	if seed==0{
		j=1
		i=0
	}else if seed==1{
			j=2
			i=-1
		}else if seed==2{
			i=0
			j=2
		}
		gen:= func() int{
			i=i+j
			return i
		}
		return gen
}

func main() {
	int_seq := sequence(0)
	odd_seq := sequence(1)
	even_seq := sequence(2)
	for i:=0;i<10;i++{
		fmt.Print(int_seq(), " ")
	}
	fmt.Println(" ")
		
	for i:=0;i<10;i++{
		fmt.Print(odd_seq(), " ")
	}
	fmt.Println(" ")


	for i:=0;i<10;i++{
		fmt.Print(even_seq(), " ")
	}
	fmt.Println(" ")
}
