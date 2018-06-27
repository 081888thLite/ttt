package game

import (
	"bufio"
	"os"
	"strings"
)

const NewLine = '\n'

type Client interface {
	Send(string)
	Receive() string
}

type TestIO struct {
	Reply    string
	Sent     string
	LastRead string
}

func (client *TestIO) Send(msg string) {
	client.Sent = msg
}

func (client *TestIO) Receive() string {
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
	received, _ := StdIn.ReadString(NewLine)
	client.LastRead = strings.Replace(received, "\n", "", -1)
	return client.LastRead
}
