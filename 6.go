package main 
import ( 
	"fmt"
	"bufio"
	"os"
	"strings"
)
func main() {
	fmt.Println("Enter a string")
	//a := make(map[int]string)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	words := strings.Split(s," ")
	wor := make(map[int]string)
	//fmt.Println(words)
	for i, word := range words{
	 	wor[i]=word	
	 } 
	
	l := len(words)
	fmt.Println(l)
	//fmt.Println(words[2])
	//fmt.Println("Before",a)
	for i:=0; i<l-1; i=i+1{
		for j:=i+1; j<l; j=j+1 {
			if wor[i]==wor[j] {
				//words = append(words[:i], words[i+1:]...)
				delete(wor,i)
			}
		} 
	}
	// a := map[string]bool
	// for _, v:= range words{
	// 	if a[words] == true {

	// 		}else{

	// 		}


	// }
	
	fmt.Println("After:",wor)
}