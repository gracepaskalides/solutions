package main

import (
	"fmt" // printing the command line 
	"os" // interop with the computer 
	"bufio" // moving io into a buffer
	"strings" // messing with text
)

// fill in this function to implement a Caesar cipher 
// shifting a string n slots to the right 
func encipher(s string, n int) string {
	// FILL ME IN TO WORK
	return "" 
}

// this function takes in a string and shift and outputs a 
// slice (a growable array) of enciphered strings
func encode(s string, n int) []string {
	var l []string  // declare a empty slice of strings
	s = strings.ToLower(s) // change to lower case
	t := strings.Split(s," ") // split a string into a slice along a space
	// here is a basic for loop in go. it calls encipher on each 
	// word in the slice t and sticks then into l 
	for i := 0; i < len(t); i++ {
		c := encipher(t[i],n) 
		l = append(l,c)
	}
	return l // when done, return the slice of encoded strings 
}

// this is the main function of the program 
func main() {
	fmt.Print("Enter your confidential phrase: ") // prints to the command line
	reader := bufio.NewReader(os.Stdin) // reads from the command line 
	// converts the useful format, aka string
	// note this could fail so it return a pair the expected string and an error
	input, err := reader.ReadString('\n') 
	// handling the error. this is the most common pattern in go :)
	if err != nil {
		fmt.Println("Something went wrong. Please try again", err)
		return
	}

	// remove the carriage return from the string
	input = strings.TrimSuffix(input, "\n")
	// applies the Caesar shift encoding with base case of 5 to the input 
	output := strings.Join(encode(input,5)[:]," ") 
	// prints back out the result
	fmt.Printf("Your encrypted phrase is %s \n",output)
}
