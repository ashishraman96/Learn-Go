package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string:")
	var count int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	words := strings.Split(s, " ")
	for _, word := range words {
		if len(word)%2 == 0 {
			count += 1
		}
	}
	fmt.Println("Number of even length words:", count)
}
