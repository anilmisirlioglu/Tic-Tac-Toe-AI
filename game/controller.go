package game

import "github.com/anilmisirlioglu/Tic-Tac-Toe-AI/math"

type Controller struct {
	game   *Game
	scores map[int]int
}

type MinimaxPacket struct {
	Vector2 math.Vector2
	Score   int
}

const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

func (c Controller) Minimax(matrix math.Matrix, depth int, player int) MinimaxPacket {
	cpu := c.game.CPU
	score := MaxInt
	if cpu.GetSymbolByInt() == player {
		score = MinInt
	}

	optimumAxisAlignment := MinimaxPacket{
		Vector2: math.Vector2{X: -1, Y: -1},
		Score:   score,
	}

	if depth == 0 || c.over() {
		return MinimaxPacket{
			Vector2: math.Vector2{X: -1, Y: -1},
			Score:   c.toScore(c.evaluate()),
		}
	}

	for row := 0; row < matrix.Rows; row++ {
		for column := 0; column < matrix.Columns; column++ {
			if matrix.GetElement(row, column) == NullIndex {
				matrix.SetElement(row, column, player)
				nextPlayer := XIndex
				if player == XIndex {
					nextPlayer = OIndex
				}

				score := c.Minimax(matrix, depth-1, nextPlayer)
				matrix.SetElement(row, column, NullIndex)

				score.Vector2 = math.Vector2{
					X: row,
					Y: column,
				}

				if player == cpu.GetSymbolByInt() {
					if score.Score > optimumAxisAlignment.Score {
						optimumAxisAlignment = score
					}
				} else {
					if score.Score < optimumAxisAlignment.Score {
						optimumAxisAlignment = score
					}
				}
			}
		}
	}

	return optimumAxisAlignment
}

func (c Controller) evaluate() int {
	matrix := c.game.matrix

	// Vertical Axis
	for v := 0; v < matrix.Rows; v++ {
		middle := matrix.GetElement(1, v)
		if middle != NullIndex {
			if matrix.GetElement(0, v) == middle && matrix.GetElement(2, v) == middle {
				return middle
			}
		}
	}

	// Horizontal Axis
	for h := 0; h < matrix.Columns; h++ {
		middle := matrix.GetElement(h, 1)
		if middle != NullIndex {
			if matrix.GetElement(h, 0) == middle && matrix.GetElement(h, 2) == middle {
				return middle
			}
		}
	}

	middle := matrix.GetElement(1, 1) // Cross Axis Middle
	if middle != NullIndex {
		// Cross Axis (Top Left, Bottom Right)
		if matrix.GetElement(0, 0) == middle && matrix.GetElement(2, 2) == middle {
			return middle
		}

		// Cross Axis (Top Right, Bottom Left)
		if matrix.GetElement(2, 0) == middle && matrix.GetElement(0, 2) == middle {
			return middle
		}
	}

	for v := 0; v < matrix.Rows; v++ {
		for h := 0; h < matrix.Columns; h++ {
			if matrix.GetElement(v, h) == NullIndex {
				return Unfinished
			}
		}
	}

	return Draw
}

func (c Controller) over() bool {
	evaluate := c.evaluate()
	return evaluate == XWon || evaluate == OWon
}

func (c Controller) toScore(winner int) int {
	score, ok := c.scores[winner]
	if !ok {
		score = 0
	}

	return score
}

func NewController(game *Game) *Controller {
	return &Controller{
		game: game,
		scores: map[int]int{
			game.CPU.GetSymbolByInt():   1,
			game.Human.GetSymbolByInt(): -1,
		},
	}
}
