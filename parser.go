package main 
import (
"fmt"
"os"
"strings"
//"regexp"
"bufio"
//"io/ioutil"
)
type funcDetail struct{
	FuncName string
	ReturnType string
	NumOfInputs int
}
func main() {
	var importCount int
	var imports []string
	function := new([]funcDetail)   
	fin,_ :=os.Open("5.go")
	scanner:=bufio.NewScanner(fin)
	for scanner.Scan(){
		line:=scanner.Text()
		if strings.Contains(line,"package"){
				tmp:= strings.Split(line," ")
				fmt.Println("Package name:", tmp[1])
		}
		if strings.Contains(line, "import"){
			importCount++
			tmp:=strings.Split(line," ")
			//fmt.Println(tmp)
			imports = append(imports,tmp[1])
		}
		if strings.Contains(line, "func"){
			cpat,_ := regexp.Compile(pat)
			result := cpat.FindAllString(s,-1)
			fmt.Println(result)

		}
	}
	fmt.Println("Number of imports:",importCount, "\t Packages imported:", imports)
}