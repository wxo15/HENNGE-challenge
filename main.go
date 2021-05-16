//Code for HENNGE challenge
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N int = 0
var X int = 0
var err error
var inp string = "/n"
var res int = 0

func main() {
	//Get standard input from users
	fmt.Scanln(&N)

	loop()
}

func loop() {
	//If N ==0, it means it had looped through N-number of test cases. Else, minus 1 from N.
	if N == 0 {
		return
	}

	//Get number of integers, X
	fmt.Scanln(&X)

	//make a slice of size X
	strslice := make([]string, X)

	//make a reader
	rd := bufio.NewReader(os.Stdin)
	inp, err = rd.ReadString('\n')
	strslice = strings.Fields(inp)

	//pass slice to add square function
	addsquares(strslice)

	//Print out results, and wipe res
	fmt.Println(res)
	res = 0

	//Recursive call
	N--
	loop()
}

func addsquares(strslice []string) {
	//declare local variable
	var intval int

	//If X ==0, it means it had looped through X-integers within the test case.
	//Else, minus 1 from X.
	if X == 0 {
		return
	}

	intval, err = strconv.Atoi(strslice[X-1])

	if intval > 0 {
		res += intval * intval
	}

	//Recursive call
	X--
	addsquares(strslice)
}
