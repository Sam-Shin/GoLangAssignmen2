package question3

import "fmt"

func MainQ3(){
	var word string = "The is a recursion call function"
	recursiveCall(word, 0)
}

func recursiveCall(s string, n int) rune{
	runes := []rune(s)
	var r rune
	if n == len(s) - 1{
		r = runes[n]
		return r
	}
	r = runes[n]
	fmt.Printf("%x\n", r)
	return recursiveCall(s, n+1)
}