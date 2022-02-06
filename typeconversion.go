package main

import (
	"fmt"
	"strconv"
)
 
func main(){
 //type conversion
//  var a int32=100
//  var b int64=int64(a)
// fmt.Println(b)

var str int  =111
var d string =strconv.Itoa(str)   //int to string

// c :=2.5
// var d float64=float64(c)

fmt.Printf("%v,%T",d,d)



}