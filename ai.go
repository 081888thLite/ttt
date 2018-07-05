package ttt

type Minimax struct {
	depth  int
	min    Piece
	max    Piece
	scores []int
	moves  []int
}

func (mm *Minimax) Score(w Piece, p Piece, depth int) int {
	switch w {
	case p:
		return 10 - depth
	case swap(p):
		return depth - 10
	default:
		return 0
	}
}

func (mm *Minimax) minimax(b Board, p Piece, depth int) int {
	blank := b.blanks()
	if len(blank) == 0 || b.wonBy() != NoOne {
		return mm.Score(b.wonBy(), p, depth)
	}
	depth += 1
	mm.scores = []int{}
	mm.moves = []int{}
	for _, move := range blank {
		nb := b[:]
		nb.Mark(move, p)
		np := swap(p)
		nmm := Minimax{}
		var result = nmm.minimax(nb, np, depth)
		mm.scores = append(mm.scores, result)
		mm.moves = append(mm.moves, move)
	}
	return mm.evaluate(b, p)
}

func (mm *Minimax) evaluate(b Board, evaluator Piece) int {
	var bestMove int
	if X == b.wonBy() {
		bestScore := -10000

		for index, score := range mm.scores {
			if score > bestScore {
				bestScore = score
				bestMove = mm.moves[index]
			}
		}
	} else {
		bestScore := 10000

		for index, score := range mm.scores {
			if score < bestScore {
				bestScore = score
				bestMove = mm.moves[index]
			}
		}
	}
	return bestMove
}

func swap(p Piece) Piece {
	if p == X {
		return O
	}
	return X
}
