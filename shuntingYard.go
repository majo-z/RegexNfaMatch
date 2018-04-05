//Shunting yard algorithm produces regular expression in postfix notation from given regular expression in infix notation
package main

import "fmt"

//inToPost takes single argument as a string and returns string
func inToPost(infix string) string{
pofix := ""
	return pofix
}

func main() {
	//Answer: ab.c*.
	fmt.Println("Infix:   ", "a.b.c*")
	fmt.Println("Postfix: ", inToPost("a.b.c*"))
	fmt.Println()

	//Answer: abd|.*
	fmt.Println("Infix:   ", "(a.(b|d)*")
	fmt.Println("Postfix: ", inToPost("(a.(b|d)*"))//postfix omits brackets
	fmt.Println()

	//Answer: abd|.c*
	fmt.Println("Infix:   ", "(a.(b|d).c*")
	fmt.Println("Postfix: ", inToPost("(a.(b|d).c*"))
	fmt.Println()

	//Answer: abb.+c.
	fmt.Println("Infix:   ", "a.(b.b)+.c*")
	fmt.Println("Postfix: ", inToPost("a.(b.b)+.c*"))
	fmt.Println()
}