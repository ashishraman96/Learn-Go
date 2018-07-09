package main 
import "fmt"
import "os/exec"
import "strings"
import "time"
func main() {

	result,_ := exec.Command("cmd","/C","ls").Output()
	s:=string(result)
	fmt.Println(s)
	entries :=strings.Split(s,"\n")
	for _,entry := range entries{
		fmt.Println(entry)
		time.Sleep(time.Second)
	}
}