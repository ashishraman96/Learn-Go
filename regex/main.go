package main 
import (
"fmt"
"regexp"
"os"
"bufio"
"strconv"
)

func main() {
	var line string
	fin,err := os.Open(os.Args[1])
			if err != nil{
					panic(err)
			}
	defer fin.Close()
	scanner:=bufio.NewScanner(fin)
	scanner.Scan()
	line=scanner.Text()
	sum(line)
	startCapital(line)
	startCapitalEndVowel(line)
	vowelInWord(line)
	allCaps(line)
	allPositiveandNegativeNumber(line)
	evenLengthNumber(line)
	customrange(line)
	validIp(line)

}

func sum(line string){
	var s int
	pat:="\\s[-]?\\d+\\s"
	cpat,_ := regexp.Compile(pat)
	result := cpat.FindAllString(line,-1)
	for _, numStr := range result{
		num,_ := strconv.Atoi(numStr)
		s=s+num 
	}
	fmt.Println("Total sum:", s)
}

func startCapital(line string){
	pat:="[A-Z]+\\w*"
	cpat,_ := regexp.Compile(pat)
	result := cpat.FindAllString(line,-1)
	fmt.Println("start capital:", result)
}

func startCapitalEndVowel(line string){
	pat:="[A-Z]+\\w*[aeiou]"
	cpat,_ := regexp.Compile(pat)
	result := cpat.FindAllString(line,-1)
	fmt.Println("start capital end vowel:", result)
}

func vowelInWord(line string){
	pat:="\\w*[aeiouAEIOU]+\\w*"
	cpat,_ := regexp.Compile(pat)
	result := cpat.FindAllString(line,-1)
	fmt.Println("vowel in word:", result)
}

func allCaps(line string){
	pat:="\\s[A-Z]+\\s"
	cpat,_ := regexp.Compile(pat)
	result := cpat.FindAllString(line,-1)
	fmt.Println("All Caps:", result)
}

func allPositiveandNegativeNumber(line string){
	pat:="\\s[-]?\\d+\\s"
	cpat,_ := regexp.Compile(pat)
	result := cpat.FindAllString(line,-1)
	fmt.Println("All positive and negative number:", result)
}

func evenLengthNumber(line string){
	pat:="\\b(\\d\\d)+\\b"
	cpat,_ := regexp.Compile(pat)
	result := cpat.FindAllString(line,-1)
	fmt.Println("even length  number:", result)
}

func customrange(line string){
	pat:=`\b(1|[1-9]\d{1}|1\d{2}|2[0-4]\d|2[0-5][0-5])\b`
	cpat,_ := regexp.Compile(pat)
	result := cpat.FindAllString(line,-1)
	fmt.Println("range from 1 to 255 number:", result)
}

func validIp(line string){
	pat:=`\b((1|[1-9]\d|1\d{2}|2[0-4]\d|2[0-5][0-5])\.){3}(1|[1-9]\d|1\d{2}|2[0-4]\d|2[0-5][0-5])\b`
	cpat,_ := regexp.Compile(pat)
	result := cpat.FindAllString(line,-1)
	fmt.Println("Valid IP:", result)


}