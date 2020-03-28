package main

import (
	"assignment2/question1"
	"assignment2/question2"
	"assignment2/question3"
	"fmt"
)

func main(){
	fmt.Print("Question 1:")
	question1.Main1()
	fmt.Println("Question 3: ")
	question3.MainQ3()
	fmt.Println("Question 2:")
	defer question2.PrintMath()
	question2.RunQ2()
}