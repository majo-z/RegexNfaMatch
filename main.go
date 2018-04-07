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
	//infix a.b.c*
	fmt.Println("Infix regular expression \"a.b.c*\" matches string \"abc\" is: ", algorithm.DoesMatch("a.b.c*", "abc"))//true

	fmt.Println("Infix regular expression \"a.b.c*\" matches string \"ab\" is: ", algorithm.DoesMatch("a.b.c*", "ab"))//true

	fmt.Println("Infix regular expression \"a.b.c*\" matches string \"abcccc\" is: ", algorithm.DoesMatch("a.b.c*", "abcccc"))//true

	fmt.Println("Infix regular expression \"a.b.c*\" matches empty string \"\" is: ", algorithm.DoesMatch("a.b.c*", ""))//false

	fmt.Println("Infix regular expression \"a.b.c*\"  matches string \"c\" is: ", algorithm.DoesMatch("a.b.c*", "c"))//false

	fmt.Println("Infix regular expression \"a.b.c*\" matches string \"abd\" is: ", algorithm.DoesMatch("a.b.c*", "abd"))//false
	fmt.Println()

	//infix a.b.c*
	fmt.Println("Infix regular expression \"a.b|c*\" matches string \"abc\" is: ", algorithm.DoesMatch("a.b|c*", "abc"))//false

	fmt.Println("Infix regular expression \"a.b|c*\" matches string \"ab\" is: ", algorithm.DoesMatch("a.b|c*", "ab"))//true

	fmt.Println("Infix regular expression \"a.b|c*\"  matches string \"ccccc\" is: ", algorithm.DoesMatch("a.b|c*", "ccccc"))//true

	fmt.Println("Infix regular expression \"a.b|c*\" matches empty string \"\" is: ", algorithm.DoesMatch("a.b|c*", ""))//true

	fmt.Println()

}
