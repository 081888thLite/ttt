package game

import (
	"bufio"
	"strings"
)

// Delimiter when reading in Strings... If you're having trouble reading
// from StdIn, and your on a Windows machine you may have to account for
// windows carriage return: '\r\n'
const ToLineEnding = '\n'

// Represents the endpoint terminal for Input and Output This could be
// extended for non-local machine use i.e. implementing a Client whose
// Write function sends an HTTPResponse to a browser
type Client interface {
	Write(string)
	Read() string
}

// StubIO is a Client Factory for ease of testing. If a unit under test
// involves behavior produced due to a client interaction, but is not
// explicitly meant to test that the interaction can occur; use StubIO to
// quickly construct a Client with deterministic values. In this way one
// can avoid having to deal with complicated spying, mocking, or memory
// intensive os.StdIn tempFile switcheroos.
type StubIO struct {
	Reply    string
	Sent     string
	LastRead string
}

// StubIO.Write simply sets the struct's Sent field to the message meant to be sent
func (client *StubIO) Write(msg string) {
	client.Sent = msg
}

// StubIO.Read sets the LastRead field (Meant to hold the last msg read in)
// to the value specified in the StubIO Client's initialization
func (client *StubIO) Read() string {
	client.LastRead = client.Reply
	return client.LastRead
}

// StdIO is the common construct for reading and writing to a client/local OS System Input/Output
type StdIO struct {
	Writer    bufio.Writer
	Reader    bufio.Reader
	lastWrite string
	lastRead  string
}

func (client *StdIO) Write(msg string) {
	client.Writer.WriteString(msg)
	client.lastWrite = msg
}

func (client *StdIO) Read(messages ...string) string {
	var received string
	if len(messages) != 0 {
		received = messages[0]
	} else {
		received, _ = client.Reader.ReadString(ToLineEnding)
	}
	client.lastRead = strings.Replace(received, "", "", -1)
	return client.lastRead
}
