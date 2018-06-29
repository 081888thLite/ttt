package ttt

import (
	"io"
	"io/ioutil"
	"testing"
)

func TestRead(t *testing.T) {
	msg := "Received\n"
	msgAfterReading := "Received"
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

	fed, _ := feed(in)
	if fed != msgAfterReading {
		t.Errorf("Expected %v,\n got: %v", msgAfterReading, fed)
	}
}

func TestSeed(t *testing.T) {
	msg := "Sent"
	out, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	_, err = io.WriteString(out, msg)
	if err != nil {
		t.Fatal(err)
	}

	_, err = out.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}
	_, err = seed(out, msg)
	if err != nil {
		t.Errorf("Expected no errors when sending '%v',\n got: %v", msg, err)
	}
}

func TestGetLastRead(t *testing.T) {
	msgStatus := MsgStatus{Msg: "Latest"}
	client := StubClient{LastRead: msgStatus}
	accessed := client.GetLastRead()
	if accessed != "Latest" {
		t.Error("Expected the client's LastRead.Msg as string,\n got:", accessed)
	}
}
func TestGetLastSent(t *testing.T) {
	msgStatus := MsgStatus{Msg: "Latest"}
	client := StubClient{LastWrote: msgStatus}
	accessed := client.GetLastSent()
	if accessed != "Latest" {
		t.Error("Expected the client's LastWrote.Msg as string,\n got:", accessed)
	}
}
