package game

import (
	"bufio"
	"os"
)

type Setup struct {
	game *Game
}

func (s Setup) Start() {
	Write("Please choose your character (X or O): ", NextLineNone, true, 2)
	for {
		reader := bufio.NewReader(os.Stdin)
		humanSymbol, _ := reader.ReadString('\n')
		if humanSymbol != "" && (humanSymbol == X || humanSymbol == O) {
			var (
				cpuSymbol string
				sequence  Player
			)

			if humanSymbol == X {
				cpuSymbol = O
			} else {
				cpuSymbol = X
			}

			human := Human{Symbol: humanSymbol}
			cpu := CPU{Symbol: cpuSymbol}

			sequence = human
			if cpuSymbol == X {
				sequence = cpu
			}

			s.game.Human = human
			s.game.CPU = cpu
			s.game.Sequence = sequence

			Write(s.game.Writer.String(), NextLineMultiply, false, 1)
			s.game.Controller = NewController(s.game)
			break
		} else {
			Write("Please choose a valid character.", NextLineMultiply, true, 2)
			Write("Please choose your character (X or O): ", NextLineNone, true, 2)
		}
	}
}

func NewSetup(game *Game) Setup {
	return Setup{
		game: game,
	}
}
