package main 
import (
"fmt"
"time"
"sync"
)
var account_state =0
var mutex = &sync.Mutex{}
func deposit(){
	for j:=0; j<1000; j++{
		mutex.Lock()
		account_state=account_state+1
		mutex.Unlock()
		time.Sleep(time.Millisecond)
	}
}

func main(){
	for i:=0; i<5; i++{
		go deposit()
	}
	time.Sleep(time.Second *5)
	fmt.Println("final balance = ", account_state)
}