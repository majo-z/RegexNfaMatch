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
			frag2.accept.edge1 = &accept //frag2 accept state is no longer accept state, it points at new created accept state

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
		} //switch

	} //for

	//if there is more than one item at the stack at the end
	if len(nfaStack) != 1 {
		fmt.Println("Uh oh", len(nfaStack), nfaStack)
	}

	return nfaStack[0] //returns nfa struct memory addresses(initial and accept state)

} //postRegToNfa

//add the state to current array, also add any states to get from that state along ϵ arrows
func addState(listOfStates []*state, single *state, accept *state) []*state {

	listOfStates = append(listOfStates, single)//pass the state

	//check if it's one of the states that has ϵ arrows coming from it
	if single != accept && single.symbol == 0 { //zero value of rune
		listOfStates = addState(listOfStates, single.edge1, accept)
		if single.edge2 != nil { //if there is a second edge
			listOfStates = addState(listOfStates, single.edge2, accept)
		}
	}
	return listOfStates
}//addState

//poMatch runs string against regular expression to evaluate if it's valid
func poMatch(poReg string, testString string) bool {

	isMatch := false             //string doesn't match the regular expression
	poNfa := postRegToNfa(poReg) //link list of states..creates NFA from regex

	current := []*state{} //keep track of list of current states in NFA
	next := []*state{}    //keep track of states that can move to from current along the arrows, including ϵ arrows

	//decide what the current array should be initially
	//add initial state of postfix nfa to current array, also add all states that can be moved to from initial state to
	current = addState(current[:], poNfa.initial, poNfa.accept) //[:] => slice

	for _, r := range testString { //loop through a character at a time
		for _, c := range current { //loop through the current state that I'm in

			if c.symbol == r {
				//add current state to next state(also add any state to get from current state along ϵ arrows)
				next = addState(next[:], c.edge1, poNfa.accept)
			}
		}
		/*
			after reading the character and following the arrows,
			the current state is old and is replaced with next state,
			next state is reset to blank set of states to get ready for next character
		*/
		current, next = next, []*state{}
	}

	//very end state(final set of states) - after reading whole string,
	for _, c := range current {
		if c == poNfa.accept { //if the current state lopping through in is equal to accept state of nfa
			isMatch = true
			break
		}
	}

	return isMatch
} //poMatch

func main() {
	//postfix regular expression test matching empty string(zero number of c's) or ab or 1/more c's
	fmt.Println(poMatch("ab.c*|", "cccc"))

}
