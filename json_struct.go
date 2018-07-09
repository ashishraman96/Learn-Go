package main 
import (
"fmt"
"encoding/json"
)

type Person struct{
	Fname string `json:"first_name"`
	Mname string `json:"middle_name"`
	Lname string `json:"last_name"`
}

func main() {
	json_string := `{"first_name":"Ashish",
						"middle_name":"Raman",
						"last_name":"Nayak"
						}`
	person := new(Person)
	json.Unmarshal([]byte(json_string), person)
	fmt.Println(*person)
	new_json, _ := json.Marshal(person)
	fmt.Printf("%s \n", new_json)
}