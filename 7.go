package main 
import (
"fmt"
"bufio"
	"os"
	"strings")
	
func main() {
	fmt.Println("Enter a string")
	//a := make(map[int]string)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	words := strings.Split(s," ")
	count := make(map[string]int, len(words))
	for _, word :=range words{
		count[word]++
	}
	
	uu:= float64(len(count))/float64(len(words))
	fmt.Print(uu)
	fmt.Println(count)

}