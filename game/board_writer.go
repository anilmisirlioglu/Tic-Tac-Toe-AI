package game

import (
	"github.com/anilmisirlioglu/Tic-Tac-Toe-AI/math"
	"strings"
)

const (
	vertical = 4

	space = " "
	wall  = "||"
)

var drawData = [][]string{
	{"#", "1", "2", "3"},
	{"a", "", "", ""},
	{"b", "", "", ""},
	{"c", "", "", ""},
}

type BoardWriter struct {
	matrix math.Matrix
	table  []string
}

func (w BoardWriter) DrawLine(line int) {
	if line < vertical {
		column := -1
		for _, row := range drawData[line] {
			if row == "" {
				item, _ := w.matrix.GetElement(line-1, column)
				row = w.detectXO(item)
			}

			w.draw([]string{space, row})

			column++
		}
	}
}

func (w BoardWriter) draw(chars []string) {
	w.table = append(w.table, chars...)
}

func (w BoardWriter) detectXO(index int) string {
	item, found := SymbolIndexes[index]
	if found {
		return NULL
	}

	return item
}

func (w BoardWriter) String() string {
	w.table = []string{}
	for i := 0; i < vertical; i++ {
		w.draw([]string{wall})
		w.DrawLine(i)
		w.draw([]string{space, wall, "\n"})
	}

	return strings.Join(w.table, "")
}

func NewBoardWriter(matrix math.Matrix) BoardWriter {
	return BoardWriter{
		matrix: matrix,
		table:  []string{},
	}
}
