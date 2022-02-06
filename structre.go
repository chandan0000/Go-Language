package main
import "fmt"
type Employee struct{
	EmpId int
	EmpName string
	EmpMobile []string

}
 func main(){
	 emp :=Employee{
		EmpId: 101,
		EmpName:  "Ravi",
		EmpMobile: []string{"1234567890","5643337689"},
	 }
emp2 :=&emp
emp.EmpName="Chandan"
fmt.Println(emp)
fmt.Println(emp2)


}

