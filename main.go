//Description: Program uses Shunting yard and Thompson's algorithms to match entered string against regular expression
//and prints either true if valid or false if invalid
//Author: Marian Ziacik
//Date: 7-apr 2018

package main

import (
	"./algorithm"
	"bufio"
	"fmt"
	"os"
)

// global scanner
var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	for { //while
		fmt.Printf("1. Enter your own experssion, 2. Print examples, 3. Exit:")
		scanner.Scan()
		userChoice := scanner.Text()

		switch userChoice {

		case "1":
			readInput()
		case "2":
			printExamples()
		case "3":
			fmt.Println("Goodbye")
			return //return back to menu option
		default:
			fmt.Println("Incorrect option, please enter 1, 2 or 3")
		}
	}

} //main

func readInput() {
	fmt.Println("Enter infix regex:")
	scanner.Scan()
	regex := scanner.Text()

	fmt.Println("Enter test string:")
	scanner.Scan()
	testString := scanner.Text()

	if algorithm.DoesMatch(regex, testString) {
		fmt.Println("the regular expression", regex, "matched the string", testString)

	} else {
		fmt.Println("the regular expression", regex, "didn't match the string", testString)

	}
	fmt.Println()
}

func printExamples() {
	fmt.Println()
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

	//infix a.(b.b)+.c*
	fmt.Println("a.(b.b)+.c* \t", algorithm.InToPost("a.(b.b)+.c*"), "\t\t\t abbc \t\t\t", algorithm.DoesMatch("a.(b.b)+.c*", "abbc"))

	fmt.Println("a.(b.b)+.c* \t", algorithm.InToPost("a.(b.b)+.c*"), "\t\t\t abb \t\t\t", algorithm.DoesMatch("a.(b.b)+.c*", "abb"))

	fmt.Println("a.(b.b)+.c* \t", algorithm.InToPost("a.(b.b)+.c*"), "\t\t\t ac \t\t\t", algorithm.DoesMatch("a.(b.b)+.c*", "ac"))
	fmt.Println()

	//infix a.(b.b)+.c*.d+
	fmt.Println("a.(b.b)+.c*.d+ \t", algorithm.InToPost("a.(b.b)+.c*.d+"), "\t\t abbcd \t\t\t", algorithm.DoesMatch("a.(b.b)+.c*.d+", "abbcd"))

	fmt.Println("a.(b.b)+.c*.d+ \t", algorithm.InToPost("a.(b.b)+.c*.d+"), "\t\t abbc \t\t\t", algorithm.DoesMatch("a.(b.b)+.c*.d+", "abbc"))

	fmt.Println("a.(b.b)+.c*.d+ \t", algorithm.InToPost("a.(b.b)+.c*.d+"), "\t\t abbcccddd \t\t", algorithm.DoesMatch("a.(b.b)+.c*.d+", "abbcccddd"))

	fmt.Println("a.(b.b)+.c*.d+ \t", algorithm.InToPost("a.(b.b)+.c*.d+"), "\t\t abbbbcccddd \t", algorithm.DoesMatch("a.(b.b)+.c*.d+", "abbbbcccddd"))
	fmt.Println()

	//infix a.(b.b)+.c?
	fmt.Println("a.(b.b)+.c? \t", algorithm.InToPost("a.(b.b)+.c?"), "\t\t\t abb \t\t\t", algorithm.DoesMatch("a.(b.b)+.c?", "abb"))

	fmt.Println("a.(b.b)+.c? \t", algorithm.InToPost("a.(b.b)+.c?"), "\t\t\t abbc \t\t\t", algorithm.DoesMatch("a.(b.b)+.c?", "abbc"))

	fmt.Println("a.(b.b)+.c? \t", algorithm.InToPost("a.(b.b)+.c?"), "\t\t\t abbcc \t\t\t", algorithm.DoesMatch("a.(b.b)+.c?", "abbcc"))
	fmt.Println()

	//infix a.b?.c+
	fmt.Println("a.b?.c+ \t\t", algorithm.InToPost("a.b?.c+"), "\t\t\t a \t\t\t\t", algorithm.DoesMatch("a.b?.c+", "a"))

	fmt.Println("a.b?.c+ \t\t", algorithm.InToPost("a.b?.c+"), "\t\t\t ac \t\t\t", algorithm.DoesMatch("a.b?.c+", "ac"))

	fmt.Println("a.b?.c+ \t\t", algorithm.InToPost("a.b?.c+"), "\t\t\t abc \t\t\t", algorithm.DoesMatch("a.b?.c+", "abc"))

	fmt.Println("a.b?.c+ \t\t", algorithm.InToPost("a.b?.c+"), "\t\t\t abbc \t\t\t", algorithm.DoesMatch("a.b?.c+", "abbc"))

	fmt.Println("a.b?.c+ \t\t", algorithm.InToPost("a.b?.c+"), "\t\t\t abccc \t\t\t", algorithm.DoesMatch("a.b?.c+", "abccc"))
	fmt.Println()

} //printExamples
