package ttt

import "fmt"

type Minimax struct {
	level    int
	min      Piece
	max      Piece
	tree     []Node
	bestMove int
}

type Node struct {
	pos   int
	score int
}

func (mm *Minimax) Score(board Board, l int, p Piece) int {
	var winner = board.wonBy()
	switch winner {
	case X:
		return 10 - l
	case swap(p):
		return l - 10
	default:
		return 0
	}
}

func (mm *Minimax) minimax(newBoard Board, p Piece, l int) int {
	blank := newBoard.blanks()
	if len(blank) == 0 || newBoard.wonBy() != NoOne {
		return mm.Score(newBoard, l, p)
	}
	l += 1
	mm.tree = []Node{}

	for _, emptySpot := range blank {
		move := Node{pos: emptySpot}
		newBoard[emptySpot] = p
		np := swap(p)
		var result = mm.minimax(newBoard[:], np, l)

		move.score = result

		newBoard[emptySpot] = Blank

		mm.tree = append(mm.tree, move)
	}
	return mm.evaluate(mm.tree, l, p)
}

func (mm *Minimax) evaluate(nodes []Node, l int, p Piece) int {
	if l == 0 {
		fmt.Println(nodes)
	}
	var bestMove int
	if p == X {
		bestScore := -100000

		for _, node := range nodes {
			if node.score > bestScore {
				bestScore = node.score
				fmt.Println("current best ", node.pos)
				bestMove = node.pos
			}
		}
	} else {
		bestScore := 100000

		for _, node := range nodes {
			if node.score < bestScore {
				bestScore = node.score
				bestMove = node.pos
			}
		}
	}
	fmt.Println(bestMove)
	fmt.Println("\nin:", nodes)
	return bestMove
}

func swap(p Piece) Piece {
	if p == X {
		return O
	}
	return X
}
