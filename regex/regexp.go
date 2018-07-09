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
	r:= cpat.MatchString(s)
	if r {
		fmt.Println("The pattern is present")
	}else{
		fmt.Println("The pattern is not present")
	}

}