/*
GoLang simple Hello World program
Author : Delveri Munang
Date :2015
File : hello.go
*/

package main

import "fmt"


func main(){
	// Go data types

	var name string = "Dario"
	word := "Some word" // implicit declaration of string

	var age int = 25
	var height float32 = 1.79
	var weight uint8 = 97

	fmt.Println("Here we go", name, word)
	fmt.Println("You are ", age , height , weight  )
}