//Shunting yard algorithm produces regular expression in postfix notation from given regular expression in infix notation
package main

import "fmt"

//inToPost converts infix regex to postfix regex
func inToPost(infix string) string {

	//map special characters into integers - order of precedence: *.|
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	//rune = character as displayed on screen (in UTF8)
	//pofix := []rune{}

	//stack used to temporarilly store operators as read from infix string
	//stack := []rune{}

	//shorthand
	pofix, stack := []rune{}, []rune{}

	return string(pofix) //cast/convert pofix(runes) to return back a string
}

func main() {
	//Answer: ab.c*.
	fmt.Println("Infix:   ", "a.b.c*")
	fmt.Println("Postfix: ", inToPost("a.b.c*"))
	fmt.Println()

	//Answer: abd|.*
	fmt.Println("Infix:   ", "(a.(b|d)*")
	fmt.Println("Postfix: ", inToPost("(a.(b|d)*")) //postfix omits brackets
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
