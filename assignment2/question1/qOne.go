package question1

import "fmt"

var WhatIsThe int = AnswerToLife()

func AnswerToLife() int{
	return 42
}
func init(){
	WhatIsThe = 0
}
func Main1(){
	if WhatIsThe == 0{
		fmt.Println("\nIt's all a lie.")
	}
}