package main

import "fmt"
import "strings"

func main(){

	str := "adgagvahba"

	fmt.Println("str[0:0]=",str[0])
	fmt.Printf("%T",string(str[0])) //unit8  str[0]
							  //string str[0:1]
								//string     string(str[0])  

	index := strings.LastIndex("badfa","f")
	fmt.Println("max=",index)

}