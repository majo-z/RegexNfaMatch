//Thomson's algorithm(construction) builds NFA from regular expression (regex is in postfix notation)
package main

import "fmt"

type state struct { //struct to store states and arrows
	symbol rune //default is 0, represents a state that has 2 arrows coming from it
	//pointers to other states - max number of arrows coming from any state is 2
	edge1 *state
	edge2 *state
}

//helper struct, keeps track of initial state and accept state of fragment of nfa
type nfa struct {
	initial *state
	accept  *state
}

//postfix regular expression to nfa
func postRegToNfa(pofix string) *nfa {
	nfaStack := []*nfa{} //pointers to nfa's
 

	return nfaStack[0] //returns nfa struct addresses
	
}

func main() {
	nfa := postRegToNfa("ab.c*|") //postfix notation
	fmt.Println(nfa)
}
