package main 
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
type matrix struct{
	row,col int
	val [][]int
}

 func main() {
 	m1 := new(matrix)
 	m1.read(os.Args[1])
 	fmt.Println(m1)
 	m2 := new(matrix)
 	m2.read(os.Args[2])
 	m3 :=new(matrix)
 	m3.add(m1,m2)
 	m3.display(os.Args[3])
}
func (m *matrix) add (m1,m2 *matrix){
		m.row=m1.row
		m.col=m1.col
		for i:=0;i<m1.row;i++{
		var col []int
			for j:=0;j<m1.col;j++{
				col=append(col,m1.val[i][j]+m2.val[i][j])
		}
		m.val = append(m.val,col)
	}
}

func (m *matrix) display(arg string){
	fmt.Println(m)
	fout,err := os.OpenFile(arg, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil{
		panic(err)
	}
	defer fout.Close()
	for i:=0; i< m.row; i++{
		for j:=0; j< m.col ; j++{
			fout.WriteString(fmt.Sprintf("%d ", m.val[i][j]))
			//
		}
		fout.Write([]byte(" "))
	}
	//fmt.Println(m.val[0][0])
}
		
func (m *matrix) read (arg string) {
	
	fin,err := os.Open(arg)
		if err != nil{
			panic(err)
		}
		defer fin.Close()
		scanner:=bufio.NewScanner(fin)
		for scanner.Scan(){
			var row []int
			line:=scanner.Text()
			m.row++
			num:=strings.Split(line," ")
			m.col = len(num)
			for _,v:=range num{
				no,_:= strconv.Atoi(v)
				row=append(row,no)
			}
			m.val = append(m.val,row)
		}
}

