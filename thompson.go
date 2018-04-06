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

	//loop through postfix regular expression rune at a time
	for _, r := range pofix {

		switch r {

		//concatenate
		case '.':
			//pointers to nfa fragments - pop off the stack
			frag2 := nfaStack[len(nfaStack)-1]    //index of last thing on nfa stack
			nfaStack = nfaStack[:len(nfaStack)-1] //pop last frag off nfa stack
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//join the fragments together in concatenation set up
			frag1.accept.edge1 = frag2.initial //first edge of accept state of frag1 points to initial state of frag2

			//push the concatenated fragment back to nfa stack
			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept}) //append new pointer to nfa struck that represents this new bigger nfa fragment

		//union(or)
		case '|':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//new state where its 2 edges point at initial states of the 2 fragments popped off the stack
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{} //new blank accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept//frag2 accept state is no longer accept state, it points at new created accept state

			//push to stack new initial & accept state
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		//kleene star
		case '*':
			//pop a fragment off the stack
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//new fragment is the old fragment with 2 extra states
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept} //new initial state points at initial state of the fragment that was popped off and new accept state

			//old except state points at its own initial state and new accept state
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			//push new fragment to the stack
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		//non special character
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			//push to stack
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}//switch

	}//for


	return nfaStack[0] //returns nfa struct memory addresses(initial and accept state)

} //postRegToNfa

func main() {
	nfa := postRegToNfa("ab.c*|") //postfix notation
	fmt.Println(nfa)
}
