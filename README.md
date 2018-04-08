# RegexToNFA - Graph Theory Project 2018
My name is Marian and I am 3rd year software development student at GMIT. 
Objective of this project was to build a non-deterministic finite automaton (NFA) from a regular expression and use the NFA to check if the regular expression matches a given string of text. The project has been split into different parts. First part was to parse the regular expression from infix to postfix notation using Shunting Yard algorithm,then to build a non-deterministic finite automaton (NFA) using Thompson's algorithm(construction) and finally to implement the matching algorithm using the NFA to test a given string to determine whether it is valid or invalid. The programs has be written in Go programming language.

## How the project was created
1. I have created Shunting yard algorithm that converts infix regular expression into postfix regular expression. The output of this algorithm is regular expression in postfix notation.
2. I have created second file, Thompson's algorithm(construction) that builds non-deterministic finite automaton (NFA) from postfix regular expression. The output of this algorithm is nfa struct initial and accept state memory addresses.
3. I have created third file, poMatch function that uses Thomson's construction to build nfa and then compares this with tested string.
The output of this function is bullion that prints either true if tested string is valid or false if it is invalid.
4. Eventually, these files were moved to algorithm package and I have created main file from where the algorithms are called. The main function consists of while loop, where the user has 3 different options: 1. to enter their regular expression in infix notation and tested string, which results in true of false statement; 2. where I have created hardcoded examples that are printed on screen - the result is grid that displays regular expression in infix notation, regular expression in postfix notation, tested string and result whether the string is valid or invalid; and 3. to exit the program.

## How to download and run the program 
The repository can be found at https://github.com/majo-z/RegexNfaMatch. 
To run the code in this repository, the files must first be compiled. The [Go compiler](https://golang.org/dl/) and [Git](https://git-scm.com/) must first be installed on your machine.
Once they are installed, the code can be compiled and run by following steps below.
We assume you are using the command line.

1. Open the Git Bash and clone the repository 
```bash
> git clone https://github.com/majo-z/RegexNfaMatch
```
2. Change into the folder.
```bash
> cd RegexNfaMatch
```
3. Compile the main.go file with the following command.
```bash
> go build main.go
```
4. Run the executable produced.
```bash
> main.exe
```
5. Alternatively, you can run the program without producing executable file as follows:
```bash
>go run main.go
```
6. To close the program
```bash
>ctrl+c
```
## References
https://swtch.com/~rsc/regexp/regexp1.html

http://jacobappleton.io/2015/07/02/regex-ii-the-shunting-yard-algorithm/#tocAnchor-1-2

https://www.cs.york.ac.uk/fp/lsa/lectures/REToC.pdf

https://autohotkey.com/docs/frame.htm#misc/RegEx-QuickRef.htm

https://regexone.com/lesson/kleene_operators

https://stackoverflow.com/questions/3075130/what-is-the-difference-between-and-regular-expressions/3075532#3075532

http://www.fon.hum.uva.nl/praat/manual/Regular_expressions_1__Special_characters.html

## Learn more about Go language
https://golang.org/

https://golang.org/doc/code.html

https://play.golang.org/


