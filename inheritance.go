package main 
import "fmt"

type person struct{
	Fname string
	Lname string
	Email string
}
type employee struct{
	person
	Eid int
	Salary int
}


func (e employee) String() string{
	return fmt.Sprintf("employee{FName: %s, Lname:%s, Email:%s}, id: %d , salary: %d", e.Fname, e.Lname, e.Email, e.Eid, e.Salary)
}

func (p person) FullName() string{
	return p.Fname+ " " + p.Lname
}

func (p person) GetEmail() string{
	return p.Email
}

func (e employee) GetID() int {
	return e.Eid
}

func (e employee) GetSalary() int {
	return e.Salary
}


func main() {
	//p:= person{"abc","xyz","abc@xyz.com"}
//	fmt.Println(p)
	//fn:= p.FullName()
//	fmt.Println(fn)
	emp:= employee{person{"xyz","abc","xyz@abc"},128,1000}
	fmt.Println(emp)
}