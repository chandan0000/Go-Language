package main

import (
	"fmt"
)
 
 func main(){//defer
defer fmt.Println("Db Connection Closing")
fmt.Println("Connection Opening.....")
fmt.Println("Database manipulation")
fmt.Println("Database manipulation")
}

