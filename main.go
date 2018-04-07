//Description: Program uses Shunting yard and Thompson's algorithms to match entered string against regular expression
//and prints either true if valid or false if invalid
//Author: Marian Ziacik
//Date: 7-apr 2018

package main

import (
	"fmt"
	"./algorithm"
)

func main() {
	//postfix regular expression test matching empty string(zero number of c's) or ab or 1/more c's
	fmt.Println("Regular expression \"ab.c*|\" matches string \"cccc\" is: ", algorithm.PoMatch("ab.c*|", "cccc"))

	fmt.Println("Regular expression \"ab.c*|\" matches string \"ab\" is: ", algorithm.PoMatch("ab.c*|", "ab"))

	fmt.Println("Regular expression \"ab.c*|\" matches string \"abc\" is: ", algorithm.PoMatch("ab.c*|", "abc"))

	fmt.Println("Regular expression \"ab.c*|\" matches empty string \"\" is: ", algorithm.PoMatch("ab.c*|", ""))

	fmt.Println("Regular expression \"ab.c*|\" matches string \"c\" is: ", algorithm.PoMatch("ab.c*|", "c"))

	fmt.Println("Regular expression \"ab.c*|\" matches string \"cccccccccc\" is: ", algorithm.PoMatch("ab.c*|", "cccccccccc"))

	fmt.Println("Regular expression \"ab.c*|\" matches string \"def\" is: ", algorithm.PoMatch("ab.c*|", "def"))


}
