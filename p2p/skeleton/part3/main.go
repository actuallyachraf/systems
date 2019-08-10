// Solution to part 3 of the Whispering Gophers code lab.
//
// This program listens on the host and port specified by the -listen flag.
// For each incoming connection, it launches a goroutine that reads and decodes
// JSON-encoded messages from the connection and prints them to standard
// output.
//
// You can test this program by running it in one terminal:
// 	$ part3 -listen=localhost:8000
// And running part2 in another terminal:
// 	$ part2 -dial=localhost:8000
// Lines typed in the second terminal should appear as JSON objects in the
// first terminal.
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
)

var listenAddr = flag.String("listen", "localhost:8000", "host:port to listen on")

type Message struct {
	Body string
}

func main() {
	flag.Parse()

	// TODO: Create a net.Listener listening from the address in the "listen" flag.
	l, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		fmt.Println(err)
	}

	for {
		// TODO: Accept a new connection from the listener.
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)

		}
		go serve(c)
	}
}

func serve(c net.Conn) {
	// TODO: Use defer to Close the connection when this function returns.
	defer c.Close()

	// TODO: Create a new json.Decoder reading from the connection.
	jsoner := json.NewDecoder(c)
	printer := json.NewEncoder(os.Stdout)
	for {
		// TODO: Create an empty message.
		var msg Message
		// TODO: Decode a new message into the variable you just created.
		err := jsoner.Decode(&msg)
		if err != nil {
			fmt.Println("failed to decode message")
			break
		}
		// TODO: Print the message to the standard output.
		err = printer.Encode(msg)
		if err != nil {
			fmt.Println("failed to print message to stdout")
			break
		}

	}
}
