package question2

import "fmt"

var x int
var y int

var sum int
var sub int
var mul int
var quo float32

func calculate(x,y int)(sum,sub,mul int,quo float32){
	sum = x + y
	sub = x - y
	mul = x * y
	quo = float32(x) / float32(y)
	return
}
func PrintMath(){
	fmt.Print("the sum is: ",sum,"\nthe difference is: ",sub,"\nthe multiplication is: ",mul,
		"\nthe quotient is: ", quo)
}

func RunQ2(){
	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d", &x)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d", &y)
	sum, sub, mul, quo = calculate(x,y)
}
