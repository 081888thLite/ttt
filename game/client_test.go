package game

import (
	"testing"
)

func TestStdIO_Read(t *testing.T) {
	stdIO := StdIO{}
	t.Log("#Read sets StdIO's lastRead field")
	userInput := "Read Works"
	stdIO.Read(userInput)
	if stdIO.lastRead != "Read Works" {
		t.Errorf("Read failed, expected 'Read Works',/n got: %v", stdIO.lastRead)
	}
}
