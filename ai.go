package ttt

import (
	"fmt"
	"log"
)

type Minimax struct {
	min Piece
	max Piece
	tree []Node
}

type Node struct {
	pos int
	score int
}

func (mm *Minimax) SetMaxMin( max Piece, min Piece ) {
	mm.max = max
	mm.min = min
}

func (mm *Minimax) Score(board Board, max Piece, min Piece) int {
	var winner = board.wonBy()
	switch winner {
	case max:
		return 10
	case min:
		return -10
	default:
		return 0
	}
}

func (mm *Minimax) minimax(newBoard Board, players [2]Player) int {
	mm.SetMaxMin(players[0].GetPiece(), players[1].GetPiece())

	openings := newBoard.blanks()

	if len(openings) == 0 {
		return mm.Score(newBoard, mm.max, mm.min)
	}
	mm.tree = []Node{}

	for _, emptySpot := range openings {
		move := Node{}
		move.pos = emptySpot

		newBoard[emptySpot] = mm.max
		fmt.Printf("Places max of %v in %v and board becomes %v\n", mm.max,emptySpot, newBoard)
		var nextPlayers = swapPlayers(players)
		var result = mm.minimax(newBoard[:], nextPlayers)

		move.score = result

		newBoard[emptySpot] = Blank

		mm.tree = append(mm.tree, move)
	}
	return mm.evaluate(mm.tree, mm.max)
}

func (mm *Minimax) evaluate(nodes []Node, max Piece) int {
	var bestMove int
	if max == mm.max {
		var bestScore = -10000

		for index, node := range nodes {
			if node.score > bestScore {
				bestScore = node.score
				bestMove = index
			}
		}
	} else {
		var bestScore = 10000

		for index, node := range nodes {
			if node.score < bestScore {
				bestScore = node.score
				bestMove = index
			}
		}
	}
	log.Println(nodes)
	return nodes[bestMove].pos;
}

func swapPlayers(players [2]Player) [2]Player {
	return [2]Player{players[1], players[0]}
}