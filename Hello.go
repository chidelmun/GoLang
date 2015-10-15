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
	var one,two,three int
	one = 1
	two = 2
	three = 3

	var a,b,c = 22,44,"word"
	fmt.Println(a,b,c)

	fmt.Println("Here we go", name, word)
	fmt.Println("You are ", age , height , weight, one,two,three  )

	fmt.Printf("%T, %T, %T\n", age, height, weight)
}