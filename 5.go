package main 
import "fmt"
import "strconv"
import "strings"


func main() {
	var n int
	fmt.Print("Enter number of IP address")
	fmt.Scan(&n)
	s := make([]string,n)
	var valid, invalid []string
	
	for i:=0;i<n;i++{
	fmt.Printf("Enter %d IP:", i+1)
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// a := scanner.Text()
	// s[i]=a
	fmt.Scan(&s[i])
	flag:=isValid(s[i])
	if flag==false{
		invalid=append(invalid,s[i])
		
	}else{
		valid=append(valid,s[i])		
	}
}
	fmt.Println("Valid:",valid, "\n", "Invalid:",invalid)	
}

func isValid(ip string) bool{
	num := strings.Split(ip, ".")
	if len(num)!=4{
		return false
	}
	for _, ipString := range num{
		ipno,_:= strconv.Atoi(ipString)
		if ipno<1 || ipno>255{
			return false
		}
	}
	return true	
}