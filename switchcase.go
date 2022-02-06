package main

import (
	"fmt"
	 
)
 
 func main(){
	//  i:=5
//  switch{
//  case i>9 &&i==4:
//       fmt.Println("hi",i)		 
//  case i==4 :
// 		 fmt.Println("Hello")


// default:
// 	  fmt.Println("None")
	 
//  }
// switch 5{
// case 5:
// 	fmt.Println("Hi")
// case 4:
// 	fmt.Println("Hello")
// default:
// 	fmt.Println("None")
// }
// switch i:=6;i {
// case 5,6,7:
// 	fmt.Println("hi")
// case 1,2,3:
// 	fmt.Println("Hello")
// case 10:
// 	fmt.Println("Hey")
// default :
//  fmt.Println("None")
// }
 
var x interface{}=1.089
switch x.(type) {
case int:
	fmt.Println("hi")
case float32:
	fmt.Println("Hello")
case  string:
	fmt.Println("Hey")
default :
 fmt.Println("None")
}

}

