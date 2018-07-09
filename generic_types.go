package main 
import (
"fmt"
"reflect"
)
func main() {
	var g interface{}
	fmt.Println("type of g = ",reflect.TypeOf(g),"value=",g)
	switch v := g.(type){
	case int:
		fmt.Println("Square of g=", v*v)
	default:
		fmt.Println(" None of the above")
	}
}