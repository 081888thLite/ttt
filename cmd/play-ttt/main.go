package main

import (
	. "github.com/081888thLite/ttt"
)

func main() {
	c := *Configure()
	g := *NewGame(c)
	g.Play()
}
