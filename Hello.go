/*
GoLang simple Hello World program
Author : Delveri Munang
Date :2015
File : hello.go
*/

package main

import "fmt"


func main(){
	var num int = 50

	if num < 100{
		fmt.Println(num < 100)
	}else if num == 100{
		fmt.Println("Numbers are equal")
	}else{
		fmt.Println("Number is greater than 100")
	}	

}