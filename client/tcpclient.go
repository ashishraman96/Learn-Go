package main 
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
	"strconv"
	)
func main() {
	conn,_ := net.Dial("tcp","127.0.0.1:8081")
	for{
		fmt.Println("Menu")
		fmt.Println("1.Testprime")
		fmt.Println("2.NextPrime")
		fmt.Println("3.Exit")
		fmt.Println("4.Download file")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter your choice")
		ch, _:= reader.ReadString('\n')
		ch= strings.Trim(ch,"\r\n")
		if ch=="3"{
			fmt.Fprintf(conn,ch+"\n")
			conn.Close()
			fmt.Println("Bye")
			break
		}
		ch1, _ := strconv.Atoi(ch)
		if ch1<1 || ch1>4{
			fmt.Println("Improper choice..")
			continue
		}
		ch = strconv.Itoa(ch1)
		fmt.Print("Enter Data:")
		data, _ := reader.ReadString('\n')
		text := ch +":" + data
		fmt.Fprintf(conn,text)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server:"+ message)
		b:= make([]byte,10)
		if ch1==4{
			fout, _ := os.OpenFile(strings.Trim(data,"\r\n"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			defer fout.Close()
			for{			
				nu,_:= bufio.NewReader(conn).Read(b)
				fout.Write(b)
				if nu==0{break}
			}
		}
	}
}
