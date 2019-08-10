// Solution to part 2 of the Whispering Gophers code lab.
//
// This program extends part 1.
//
// It makes a connection the host and port specified by the -dial flag, reads
// lines from standard input and writes JSON-encoded messages to the network
// connection.
//
// You can test this program by installing and running the dump program:
// 	$ go get github.com/campoy/whispering-gophers/util/dump
// 	$ dump -listen=localhost:8000
// And in another terminal session, run this program:
// 	$ part2 -dial=localhost:8000
// Lines typed in the second terminal should appear as JSON objects in the
// first terminal.
//
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var dialAddr = flag.String("dial", "localhost:8000", "host:port to dial")
var message = flag.String("message", "Hello", "what to tell the guy")
var duration = flag.Duration("delay", 2*time.Second, "how long do we wait")

type Message struct {
	Body string
}

func main() {
	// TODO: Parse the flags.
	flag.Parse()
	// TODO: Open a new connection using the value of the "dial" flag.
	conn, err := net.Dial("tcp", *dialAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	// TODO: Don't forget to check the error.
	s := bufio.NewScanner(os.Stdin)
	e := json.NewEncoder(conn)
	// TODO: Create a json.Encoder writing into the connection you created before.
	for s.Scan() {
		if s.Text() == "close" {
			conn.Close()
			return
		}
		m := Message{Body: s.Text()}
		err := e.Encode(m)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
