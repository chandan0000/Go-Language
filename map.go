package main
import "fmt"
 func main(){
//  empSal :=make(map[string]int)  //declartion
 empSal :=map[string]int{  //declartion with inilatzation
	 "neha":20000,
	 "raj":25000,
	 "Atul":15000,
 }
//  empSal["Atul"]=18000
//  empSal["chandaqn"]=23000
//  fmt.Println(empSal["Atul"])
//  fmt.Println(empSal)
//  emp := empSal //ponting address 
//  emp["Rohan"]=1111111
//  println(emp)
// delete(empSal,"Raj")
// fmt.Println(empSal)
fmt.Println(empSal)
// ankit,ok  :=empSal["Ankit"]

// var _ int 
_,flag :=empSal["Neha"]
fmt.Println(flag)

}