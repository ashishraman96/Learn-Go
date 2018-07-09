package main 
import (
	"fmt"
	"time"
)

type Ball struct{
	hits int
}

func player(name string, table chan Ball){
	for{
	b:=  <-table
	b.hits = b.hits+1
	fmt.Println(name,"Hits:", b.hits)
	time.Sleep(time.Second)
	table<-b }
}

func main() {
	ball :=Ball{}
	table := make(chan Ball)
	go player("one",table)
	go player("two",table)
	table <- ball
	time.Sleep(time.Second*60)
	<-table
}