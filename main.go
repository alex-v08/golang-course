package main

import "fmt"

func main() {
	var myIntVar int 
	myIntVar = -12

	var myUintVar uint
	myUintVar = 12 // Change the value to a positive number
	fmt.Println("My variable is:", myIntVar, myUintVar)

	

	var myStringVar string
	myStringVar = "Hello, World!" // Initialize myStringVar with a value
	fmt.Println("my variable is ", myStringVar)

fmt.Println("my variable addres is ", &myStringVar)



}

