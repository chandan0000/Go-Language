package main

import "fmt"
 
func main(){

var obj =Bike{
	Name: "Honda",
	Colour: "Black",
	Price: 100000,
}
obj.ShowDetails()
fmt.Println( obj.ShowName())
}


type vehicle interface{
	ShowDetails()
	ShowName()string
}
type Bike struct{
Name,Colour string
Price float64	
}
func (bike Bike) ShowDetails(){
	fmt.Println("Bike Name : ",bike.Name)
	fmt.Println("Bike Name : ",bike.Colour)
	fmt.Println("Bike Name : ",bike.Price)
}
 func (bike Bike)ShowName()string{
	 return bike.Name
 }