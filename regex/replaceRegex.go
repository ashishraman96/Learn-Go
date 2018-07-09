package main 
import (
"fmt"
"regexp"
)
func main() {
	s:= "hello hi 1234 or 56789 done"
	var pat string
	fmt.Println("Enter a pattern")
	fmt.Scan(&pat)
	cpat,_ := regexp.Compile(pat)
	s1 := cpat.ReplaceAllString(s, "ABC")
	fmt.Println(s1)
}