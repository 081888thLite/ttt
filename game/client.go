package game

import (
	"bufio"
	"os"
	"strings"
)

const ToLineEnding = '\n'

type Client interface {
	Send(string)
	Receive() string
}

type StubIO struct {
	Reply    string
	Sent     string
	LastRead string
}

func (client *StubIO) Send(msg string) {
	client.Sent = msg
}

func (client *StubIO) Receive() string {
	client.LastRead = client.Reply
	return client.LastRead
}

type StdIO struct {
	LastRead string
}

func (client *StdIO) Send(msg string) {
	StdOut := bufio.NewWriter(os.Stdin)
	StdOut.WriteString(msg)
}

func (client *StdIO) Receive() string {
	StdIn := bufio.NewReader(os.Stdin)
	received, _ := StdIn.ReadString(ToLineEnding)
	client.LastRead = strings.Replace(received, "", "", -1)
	return client.LastRead
}
