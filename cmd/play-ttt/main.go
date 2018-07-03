package main

import (
	. "github.com/081888thLite/ttt"
	"fmt"
)

func main() {
	c := *Configure()
	g := *NewGame(c)
	fmt.Printf("%T with players of type %T & %T", c, c.Players[0], c.Players[1])
	g.Play()
}
