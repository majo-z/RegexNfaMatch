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
	fmt.Println("Infix regex: a.b.c*; postfix regex:", algorithm.InToPost("a.b.c*") + "; test string: abc; match:", algorithm.DoesMatch("a.b.c*", "abc"))//true

	fmt.Println("Infix regex: a.b.c*; postfix regex:", algorithm.InToPost("a.b.c*") + "; test string: ab; match:", algorithm.DoesMatch("a.b.c*", "ab"))//true

	fmt.Println("Infix regex: a.b.c*; postfix regex:", algorithm.InToPost("a.b.c*") + "; test string: abcccc; match:", algorithm.DoesMatch("a.b.c*", "abcccc"))//true

	fmt.Println("Infix regex: a.b.c*; postfix regex:", algorithm.InToPost("a.b.c*") + "; test string: \"empty string\"; match:", algorithm.DoesMatch("a.b.c*", " "))//false

	fmt.Println("Infix regex: a.b.c*; postfix regex:", algorithm.InToPost("a.b.c*") + "; test string: c; match:", algorithm.DoesMatch("a.b.c*", "c"))//false

	fmt.Println("Infix regex: a.b.c*; postfix regex:", algorithm.InToPost("a.b.c*") + "; test string: abd; match:", algorithm.DoesMatch("a.b.c*", "abd"))//false
	fmt.Println()

	//infix a.b.c*
	fmt.Println("Infix regex: a.b|c*; postfix regex:", algorithm.InToPost("a.b|c*") + "; test string: abc; match:", algorithm.DoesMatch("a.b|c*", "abc"))//false

	fmt.Println("Infix regex: a.b|c*; postfix regex:", algorithm.InToPost("a.b|c*") + "; test string: ab; match:", algorithm.DoesMatch("a.b|c*", "ab"))//true

	fmt.Println("Infix regex: a.b|c*; postfix regex:", algorithm.InToPost("a.b|c*") + "; test string: ccccc; match:", algorithm.DoesMatch("a.b|c*", "ccccc"))//true

	fmt.Println("Infix regex: a.b|c*; postfix regex:", algorithm.InToPost("a.b|c*") + "; test string: c; match:", algorithm.DoesMatch("a.b|c*", "c"))//true

	fmt.Println("Infix regex: a.b|c*; postfix regex:", algorithm.InToPost("a.b|c*") + "; test string: \"empty string\"; match:", algorithm.DoesMatch("a.b|c*", ""))//true
	fmt.Println()


}
