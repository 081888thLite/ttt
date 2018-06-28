package game

import (
	"io"
	"io/ioutil"
	"testing"
)

func TestInput(t *testing.T) {
	var (
		userInput string
	)
	msg := "Received\n"
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, msg)
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}

	userInput = Read(in)
	if userInput != "Received" {
		t.Errorf("Expected %v,\n got: %v", msg, userInput)
	}
}
