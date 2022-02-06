package main

import (
	"fmt"
)
 
type Employee struct{
	EmpId int
 
}
 func main(){ 
	//  var x int =90
	//  var b *int=&x
	//  fmt.Println(x)
	//  fmt.Println(&x)
	//  fmt.Println(b)
	//  fmt.Println(*b)
  var e Employee
	var emp *Employee
	fmt.Println(emp)
	emp=new(Employee)
	fmt.Println(*emp)
    fmt.Println(e)
 
}

