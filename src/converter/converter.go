package main

import (
	"bufio"
	"converter/roman"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Declare a command rompt reader.
	reader := bufio.NewReader(os.Stdin)
	// Ask the user to type in a number.
	fmt.Print("Enter number to convert: ")
	// Retrieve the typed number upon '\n' being hit.
	text, _ := reader.ReadString('\n')
	// Convert the number retrieved as a string, into an int.
	number, _ := strconv.Atoi(text)
	// Print the result of the roman conversion of the specified number.
	fmt.Println(roman.Convert(number))
}
