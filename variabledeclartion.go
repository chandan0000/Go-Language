// package main
// import "fmt"

// func main(){
// 	//1st method
// 	var a int //declaration
// 	a=55 //initialization
// 	//2nd method
// 	var b int=500//declartion with initalaztion
// 	var c=100 //declartion without type

//     //3rd method
// 	d :=1111

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }



 
// //global variable --(Pascal Case) 
// var Val2 int =100 
// //package level -- (Camel Case)
// var myValue int=500

// //Shadowing
// var val1 int =100
// func main(){
// 	//local

//  var  val1=55 //local variable}


package main
import "fmt"
 
func main(){
	//local
	var val=100
	fmt.Printf("%v,%T",val,val)


}