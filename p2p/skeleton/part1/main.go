// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//	Hello!
// Produces this output:
//	{"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Body string
}

func main() {
	// TODO: Create a new bufio.Scanner reading from the standard input.
	reader := bufio.NewScanner(os.Stdin)
	// TODO: Create a new json.Encoder writing into the standard output.
	jsoner := json.NewEncoder(os.Stdout)
	for reader.Scan() {
		// TODO: Create a new message with the read text.
		msg := Message{Body: reader.Text()}
		err := jsoner.Encode(msg)
		// TODO: Encode the message, and check for errors!
		if err != nil {
			fmt.Println("failed to encoder to json :", err)
		}
	}
	// TODO: Check for a scan error.
}
