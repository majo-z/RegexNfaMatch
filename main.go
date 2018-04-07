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

	fmt.Println("Infix regex \t Postfix regex \t\t Test string \t Match")
	fmt.Println("===========================================================")

	//infix a.b.c*
	fmt.Println("a.b.c* \t\t\t", algorithm.InToPost("a.b.c*"), "\t\t\t abc \t\t\t", algorithm.DoesMatch("a.b.c*", "abc"))

	fmt.Println("a.b.c* \t\t\t", algorithm.InToPost("a.b.c*"), "\t\t\t ab \t\t\t", algorithm.DoesMatch("a.b.c*", "ab"))

	fmt.Println("a.b.c* \t\t\t", algorithm.InToPost("a.b.c*"), "\t\t\t abcccc \t\t", algorithm.DoesMatch("a.b.c*", "abcccc"))

	fmt.Println("a.b.c* \t\t\t", algorithm.InToPost("a.b.c*"), "\t\t\t empty string \t", algorithm.DoesMatch("a.b.c*", ""))

	fmt.Println("a.b.c* \t\t\t", algorithm.InToPost("a.b.c*"), "\t\t\t c \t\t\t\t", algorithm.DoesMatch("a.b.c*", "c"))

	fmt.Println("a.b.c* \t\t\t", algorithm.InToPost("a.b.c*"), "\t\t\t abd \t\t\t", algorithm.DoesMatch("a.b.c*", "abd"))
	fmt.Println()


	//infix a.b|c*
	fmt.Println("a.b|c* \t\t\t", algorithm.InToPost("a.b|c*"), "\t\t\t abc \t\t\t", algorithm.DoesMatch("a.b|c*", "abc"))

	fmt.Println("a.b|c* \t\t\t", algorithm.InToPost("a.b|c*"), "\t\t\t ab \t\t\t", algorithm.DoesMatch("a.b|c*", "ab"))

	fmt.Println("a.b|c* \t\t\t", algorithm.InToPost("a.b|c*"), "\t\t\t ccccc \t\t\t", algorithm.DoesMatch("a.b|c*", "ccccc"))

	fmt.Println("a.b|c* \t\t\t", algorithm.InToPost("a.b|c*"), "\t\t\t empty string \t", algorithm.DoesMatch("a.b|c*", "c"))

	fmt.Println("a.b|c* \t\t\t", algorithm.InToPost("a.b|c*"), "\t\t\t c \t\t\t\t", algorithm.DoesMatch("a.b|c*", "c"))
	fmt.Println()


	//infix a.(b|d).c*
	fmt.Println("a.(b|d).c* \t\t", algorithm.InToPost("a.(b|d).c*"), "\t\t\t abc \t\t\t", algorithm.DoesMatch("a.(b|d).c*", "abc"))

	fmt.Println("a.(b|d).c* \t\t", algorithm.InToPost("a.(b|d).c*"), "\t\t\t ab \t\t\t", algorithm.DoesMatch("a.(b|d).c*", "ab"))

	fmt.Println("a.(b|d).c* \t\t", algorithm.InToPost("a.(b|d).c*"), "\t\t\t adc \t\t\t", algorithm.DoesMatch("a.(b|d).c*", "adc"))

	fmt.Println("a.(b|d).c* \t\t", algorithm.InToPost("a.(b|d).c*"), "\t\t\t adcccc \t\t", algorithm.DoesMatch("a.(b|d).c*", "adcccc"))

	fmt.Println("a.(b|d).c* \t\t", algorithm.InToPost("a.(b|d).c*"), "\t\t\t abdc \t\t\t", algorithm.DoesMatch("a.(b|d).c*", "abdc"))
	fmt.Println()


}
