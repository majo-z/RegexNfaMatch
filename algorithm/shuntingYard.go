//Shunting yard algorithm produces regular expression in postfix notation from given regular expression in infix notation
package algorithm

import "fmt"

//inToPost converts infix regex to postfix regex
func inToPost(infix string) string {

	//map special characters into integers - order of precedence: *.|
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	//stack used to temporarily store operators as read from infix string
	pofix, stack := []rune{}, []rune{}

	//loop over the infix(input) string to convert it to postfix string
	for _, r := range infix { //r is char, range converts the array of strings to runes

		switch {

		case r == '(':
			stack = append(stack, r) //temporarily append bracket at the end of stack

		case r == ')':

			//loop over the stack until open bracket is found
			for stack[len(stack)-1] != '(' {

				pofix = append(pofix, stack[len(stack)-1]) //pop the chars off the stack and append on to postfix(output)
				stack = stack[:len(stack)-1] // colon means up to last character
			}
			stack = stack[:len(stack)-1] //pop bracket "(" off the end of the stack

		case specials[r] > 0: //special characters...0 value returned back if special character not in the map

			//while the stack still has elements and precedence of the current character <= to the precedence of the top element of the stack
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				pofix, stack = append(pofix, stack[len(stack)-1]), stack[:len(stack)-1]
			}
			stack = append(stack, r)//element of the stack has less precedence

		default: //r is neither a bracket nor a special character(may be a, b, c, d...)
			pofix = append(pofix, r)

		} //switch

	} //for

	//if there is anything on the stack
	for len(stack) > 0 {
		pofix, stack = append(pofix, stack[len(stack)-1]), stack[:len(stack)-1]
	}

	return string(pofix) //cast/convert pofix(runes) to return back a string
}

func main() {
	//Answer: ab.c*.
	fmt.Println("Infix:   ", "a.b.c*")
	fmt.Println("Postfix: ", inToPost("a.b.c*"))
	fmt.Println()

	//Answer: abd|.*
	fmt.Println("Infix:   ", "(a.(b|d))*")
	fmt.Println("Postfix: ", inToPost("(a.(b|d))*")) //postfix omits brackets
	fmt.Println()

	//Answer: abd|.c*
	fmt.Println("Infix:   ", "a.(b|d).c*")
	fmt.Println("Postfix: ", inToPost("a.(b|d).c*"))
	fmt.Println()

	//Answer: abb.+c.
	fmt.Println("Infix:   ", "a.(b.b)+.c*")
	fmt.Println("Postfix: ", inToPost("a.(b.b)+.c*"))
	fmt.Println()
}
