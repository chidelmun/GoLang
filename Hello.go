/*
GoLang simple Hello World program
Author : Delveri Munang
Date :2015
File : hello.go
*/

package main

import "fmt"

func max(a,b int) int {
	if a > b{
		return a
	}else{
		return b
	}
}

func main(){
fmt.Println(max(4,5))

}