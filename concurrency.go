package main 
import (
		"fmt"
		"time"
		)

func sayHi(){
	for i:=0; i<10; i=i+1{
		fmt.Println("Say Hi...", i)
		time.Sleep(time.Second)
	}
}
func sayHello(){
	for i:=0; i<10; i=i+1{
		fmt.Println("Say Hello...", i)
		time.Sleep(time.Second)
	}
}
func main() {
	go sayHi()
	go sayHello()
	for k:=0; k<10; k=k+1{
		fmt.Println("Main...", k)
		time.Sleep(time.Second)
	}	
}
