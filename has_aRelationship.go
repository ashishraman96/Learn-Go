package main 
import "fmt"

type person struct{
	Fname string
	Lname string
	Email string
}
func (p person) FullName() string{
	return p.Fname+ " " + p.Lname
}

type employee struct{
	Person person
	Eid int
	Salary int
}


func (e employee) String() string{
	return fmt.Sprintf("employee{FName: %s, Lname:%s, Email:%s}, id: %d , salary: %d", e.Person.Fname, e.Person.Lname, e.Person.Email, e.Eid, e.Salary)
}

func (e employee) FullName() string{
	return e.Person.FullName()
}


// func (p person) GetEmail() string{
// 	return p.Email
// }

// func (e employee) GetID() int {
// 	return e.Eid
// }

// func (e employee) GetSalary() int {
// 	return e.Salary
// }


func main() {
	//p:= person{"abc","xyz","abc@xyz.com"}
//	fmt.Println(p)
	//fn:= p.FullName()
//	fmt.Println(fn)
	p:=person{"xyz","abc","xyz@abc"}
	emp:= employee{p,128,1000}
	fmt.Println(emp.FullName())
}