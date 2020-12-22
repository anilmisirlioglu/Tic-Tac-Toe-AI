package game

import (
	"bufio"
	"fmt"
	"github.com/anilmisirlioglu/Tic-Tac-Toe-AI/math"
	"math/rand"
	"os"
	"strings"
	"time"
)

const maxRound = 9

type Game struct {
	currentRound      int
	isFinish          bool
	Human             Human
	CPU               CPU
	Writer            BoardWriter
	matrix            math.Matrix
	Controller        *Controller
	axisAlignmentData map[string]math.Vector2
	Sequence          Player
}

func (g Game) Start() {
	Write("\n", NextLineNone, false, 0)
	Write("Invincible Tic Tac Toe AI", NextLineMultiply, true, 2)

	setup := NewSetup(&g)
	setup.Start()

	for g.currentRound < maxRound && !g.isFinish {
		if g.Sequence.GetName() == "Human" {
			Write(fmt.Sprintf("Player's turn (%s):", g.Human.Symbol), NextLineNone, true, 2)

			reader := bufio.NewReader(os.Stdin)
			line, _ := reader.ReadString('\n')

			axis := g.GetAxisAlignment(strings.ToUpper(line))
			if axis == nil {
				Write("Invalid field. Please select an area on the game table.", NextLineMultiply, true, 2)
			} else {
				element := g.matrix.GetElement(axis.Y, axis.Y)
				if element == NullIndex {
					g.next(*axis, g.CPU)
				} else {
					Write("This field is occupied. Please select an empty space.", NextLineMultiply, true, 2)
				}
			}
		} else {
			var axis math.Vector2
			if g.currentRound == 0 {
				keys := make([]string, 0, len(g.axisAlignmentData))
				for key := range g.axisAlignmentData {
					keys = append(keys, key)
				}

				rand.Seed(time.Now().Unix())
				n := rand.Int() % len(keys)
				axis = *g.GetAxisAlignment(keys[n])
			} else {
				axis = g.
					Controller.
					Minimax(g.matrix, maxRound-g.currentRound, g.CPU.GetSymbolByInt()).
					Vector2

				Write(fmt.Sprintf("CPU's turn (%s): %s\n", g.CPU.Symbol, axis.String()), NextLineNone, true, 2)
				g.next(axis, g.Human)
			}
		}
	}
}

func (g Game) next(axis math.Vector2, sequence Player) {
	g.matrix.SetElement(axis.X, axis.Y, sequence.GetSymbolByInt())

	Write(g.Writer.String(), NextLineMultiply, false, 1)

	isFinish := g.Controller.evaluate()
	if isFinish != Unfinished {
		g.isFinish = true
		var winner string
		switch isFinish {
		case XWon:
			winner = fmt.Sprintf("X Won (%s)", g.Sequence.GetName())
		case OWon:
			winner = fmt.Sprintf("Y Won (%s)", g.Sequence.GetName())
		default:
			winner = "Draw Over"
		}

		Write(fmt.Sprintf("Game over. Result: {%s}", winner), NextLineBehind, true, 2)
	}

	g.Sequence = sequence
	g.currentRound++
}

func (g Game) GetAxisAlignment(axis string) *math.Vector2 {
	item, ok := g.axisAlignmentData[axis]
	if !ok {
		return nil
	}

	return &item
}

func NewGame() *Game {
	axisAlignmentData := map[string]math.Vector2{
		"A1": {X: 0, Y: 0},
		"A2": {X: 0, Y: 1},
		"A3": {X: 0, Y: 2},
		"B1": {X: 1, Y: 0},
		"B2": {X: 1, Y: 1},
		"B3": {X: 1, Y: 2},
		"C1": {X: 2, Y: 0},
		"C2": {X: 2, Y: 1},
		"C3": {X: 2, Y: 2},
	}
	matrix := math.NewMatrix(Vertical, Horizontal, nil)

	return &Game{
		currentRound:      0,
		isFinish:          false,
		Human:             nil,
		CPU:               nil,
		Writer:            NewBoardWriter(&matrix),
		matrix:            matrix,
		axisAlignmentData: axisAlignmentData,
		Sequence:          nil,
	}
}
