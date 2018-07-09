package main 
import (
"math"
	"fmt"
	"net"
	"bufio"
	"strings"
	"os"
	"strconv"
	)

func checkPrime(n int) bool{
	flag:=true
	sq:= math.Floor(math.Sqrt(float64(n)))
	//fmt.Println("number:", n, "root:",sq)
	for i:=2; i <= int(sq) ; i++{
		if n%i==0{
			flag=false
			return flag
		}
	
	}
	return flag
}

func nextPrime(n int) int {
	for{
		n=n+1
		if checkPrime(n){return n}
	}
}

func handler(conn net.Conn){
	for{
		message,_ := bufio.NewReader(conn).ReadString('\n')
		if message=="3\n"{
			conn.Close()
			fmt.Println("connection closed for :", conn.RemoteAddr())
			return
		}
		L:= strings.Split(message,":")
		var newmessage string
		if L[0] == "1"{
			num,_ := strconv.Atoi(strings.Trim(L[1],"\r\n"))
			result:= checkPrime(num)
			newmessage= strconv.FormatBool(result) + "\n"
		}
		if L[0]=="2"{
			num,_:= strconv.Atoi(strings.Trim(L[1],"\r\n"))
			result:= nextPrime(num)
			newmessage=strconv.Itoa(result) + "\n"
		}
		conn.Write([]byte(newmessage + "\n"))
		if L[0]=="4"{
			fin,_ := os.Open(strings.Trim(L[1],"\r\n"))
			b:=make([]byte,10)
			for{
				n,_ :=fin.Read(b) //n is the number of bytes read
				if n==0 {break} //end of file
				conn.Write(b)
			}
		}
		
	}
}

func main(){
	fmt.Println("server is up")
	listener, _ :=net.Listen("tcp",":8081")
	for{
		conn, _ := listener.Accept()
		fmt.Println("Connection estabilished from :", conn . RemoteAddr())
		go handler(conn)
	}
}
