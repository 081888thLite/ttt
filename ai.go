package ttt

type Minimax struct {
	caller Piece
	min    Piece
	max    Piece
	tree   []Node
}

type Node struct {
	pos   int
	score int
}

func (mm *Minimax) SetCaller(comp Hard) {
	mm.caller = comp.Piece
}

func (mm *Minimax) Score(board Board) int {

	var winner= board.wonBy()
	//log.Printf("Which was a win, making the WINNER %v\n", winner)
	switch winner {
	case mm.caller:
		//log.Printf("Me SO I GAVE 10")
		return 10
	case NoOne:
		//log.Printf("NOONE SO I GAVE 0")
		return 0
	default:
		//log.Printf("CALLER SO I GAVE 10")
		return -10
	}
}

func (mm *Minimax) minimax(newBoard Board, players [2]Player) int {
	mm.min = players[0].GetPiece()
	mm.max = players[1].GetPiece()

	openings := newBoard.blanks()

	if len(openings) == 0 || newBoard.wonBy() != NoOne {
		return mm.Score(newBoard)
	}
	mm.tree = []Node{}

	for _, emptySpot := range openings {
		move := Node{}
		move.pos = emptySpot

		newBoard[emptySpot] = mm.max
		//fmt.Printf("Places max of %v in %v and board becomes %v\n", mm.max, emptySpot, newBoard)
		var nextPlayers = swapPlayers(players)
		var result = mm.minimax(newBoard[:], nextPlayers)

		move.score = result

		newBoard[emptySpot] = Blank

		mm.tree = append(mm.tree, move)
	}
	return mm.evaluate(mm.tree, mm.max)
}

func (mm *Minimax) evaluate(nodes []Node, evaluator Piece) int {
	var bestMove int
	if evaluator == mm.caller {
		bestScore := -10000

		for index, node := range nodes {
			if node.score > bestScore {
				bestScore = node.score
				bestMove = index
			}
		}
	} else {
		bestScore := 10000

		for index, node := range nodes {
			if node.score < bestScore {
				bestScore = node.score
				bestMove = index
			}
		}
	}
	return nodes[bestMove].pos
}

func swapPlayers(players [2]Player) [2]Player {
	return [2]Player{players[1], players[0]}
}
