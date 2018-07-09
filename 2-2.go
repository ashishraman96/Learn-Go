package main 
import (
	"fmt"
	"os"
	"strconv"
	)
func main() {
	n:= os.Args[1]
	fin,err := os.Open(os.Args[2])
		if err != nil{
			panic(err)
		}
		defer fin.Close()
	fs,_ := fin.Stat()
	length := fs.Size()
	fmt.Print(length)
	t,_:=strconv.Atoi(n)
	segL := length/int64(t)
	//r := length%int64(t)
	b:= make([]byte,segL)
	for i:=1;i<t;i++{
		fout,err := os.Create("temp" +  strconv.FormatInt(int64(i), 10) + ".txt")
			if err != nil{
				panic(err)
			}
			defer fout.Close()
			
			fin.Read(b)
			fout.Write(b)
		
	}
	fin.Seek(segL*int64((t-1)),0)
	fout2,err := os.Create("temp" +n + ".txt")
						if err != nil{
								panic(err)
						}
	for {
					n,_ := fin.Read(b)
					if n==0 {break}
					fout2.Write(b)
					defer fout2.Close()
		}
			
		
			
}
	
