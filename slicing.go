package main
import "fmt"
 func main(){
//  var arr [4] int =[4]int {1,2,3,4} //array 
// 2nd method
// var arr =[...]int {1,2,3,4} //array 

//slicing 
// var arr[]int=[]int{1,2,3,4}
// arr2 :=arr //address copy
// arr[1]=200
// fmt.Println(arr)
// fmt.Println(cap(arr))
//  fmt.Println(arr2)

// 3nd method
var arr [] int= []int {1,2,3,4}
arr2 :=append(arr,100)

arr3 :=append(arr,arr2...)

fmt.Println(arr3)
// fmt.Println(len(arr))
 

}