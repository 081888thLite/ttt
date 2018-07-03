package ttt

import (
	"fmt"
	"os"
)

type MsgStatus struct {
	Msg string
	Err error
}

func feed(source *os.File) (string, error) {
	var msg string
	if source == nil {
		source = os.Stdin
	}
	_, err := fmt.Fscanf(source, "%v", &msg)
	return msg, err
}

func seed(source *os.File, msg string) (string, error) {
	if source == nil {
		source = os.Stdout
	}
	_, err := fmt.Fprintf(source, "%v", msg)
	return msg, err
}

type Sys struct {
	LastWrote MsgStatus
	LastRead  MsgStatus
}

func (client *Sys) Write(msg string) {
	seed(os.Stdout, msg)
}

func (client *Sys) Read() {
	fed, err := feed(os.Stdin)
	client.LastRead = MsgStatus{fed, err}
}

func (client *Sys) GetLastSent() string {
	return client.LastWrote.Msg
}

func (client *Sys) GetLastRead() string {
	return client.LastRead.Msg
}

type StubClient struct {
	LastWrote MsgStatus
	LastRead  MsgStatus
}

func (client *StubClient) Write(msg string) {
	client.LastWrote = MsgStatus{msg, nil}
}
func (client *StubClient) Read() {}

func (client *StubClient) GetLastSent() string {
	return client.LastWrote.Msg
}

func (client *StubClient) GetLastRead() string {
	return client.LastRead.Msg
}
