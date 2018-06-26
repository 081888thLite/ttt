package main

import (
	"fmt"
	. "github.com/081888thLite/ttt"
)

func main() {
	c := *Configure()
	fmt.Printf("%T with players of type %T & %T", c, c.Players[0], c.Players[1])
	g := *NewGame(c)
	g.Play()
}
