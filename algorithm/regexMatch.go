package algorithm

//poMatch runs string against regular expression to evaluate if it's valid
func PoMatch(poReg string, testString string) bool {//change poMatch to PoMatch to be seen outside of the package

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

//build postfix regular expression from infix notation...Shunting yard algorithm(inToPost function)
func DoesMatch(infix string, testString string) bool {
	postFix := InToPost(infix) // convert infix into postfix
	return PoMatch(postFix, testString)

}

