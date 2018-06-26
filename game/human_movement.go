package game

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const PromptForMove = "To make a move, enter a number corresponding to an open board position and press 'return':"

var Input = bufio.NewReader(os.Stdin)
var Output = bufio.NewWriter(os.Stdout)
//TODO: var Output = bufio.NewWriter(default to os.Stdout, but make this configurable for testing to capture output)

type Human struct {}

func (human Human) SetInput(reader bufio.Reader) {
	*Input = reader
}

func (human Human) SetOutput(writer bufio.Writer) {
	*Output = writer
}

func (human Human) getInput() string {
	fmt.Print("-> ")
	received, _ := Input.ReadString('\n')
	text := strings.Replace(received, "\n", "", -1)
	return text
}

func (human Human) getMove(board Board) int {
	Output.WriteString(PromptForMove)
	entered := human.getInput()
	move, _ := strconv.Atoi(entered)
	return move
}
