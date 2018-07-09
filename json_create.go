package main 
import (
"fmt"
"encoding/json"
"os"
"io/ioutil"
)
type name struct{
	Fname string
	Lname string
}

type person struct{
	Name name
	Eid int
	Phone string
}

func main() {
	p:= person{name{"Ashish","Nayak"}, 605185, "1234567890"}
	fmt.Println(p)

	json_p ,_ := json.Marshal(p)
	fout, _ := os.Create("records.json")
	fout.Write(json_p)
	fout.Close()
	
	p1 := person{}
	b,_ := ioutil.ReadFile("records.json")
	json.Unmarshal(b, &p1)
	fmt.Println(p1)
}