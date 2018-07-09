package main 
import (
		"fmt"
		"os"
)
func main() {
	l:=len(os.Args)
	fmt.Print(l)
	fout,err := os.OpenFile(os.Args[l-1], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil{
					panic(err)
			}
	for i:=1;i<l-1;i++{
		fin,err := os.Open(os.Args[i])
		if err != nil{
			panic(err)
		}
		defer fin.Close()
		defer fout.Close()
		b:=make([]byte,1)
		for{
				n,_ :=fin.Read(b) //n is the number of bytes read
				if n==0 {break} //end of file
				fout.Write(b)
		}
	}

}
