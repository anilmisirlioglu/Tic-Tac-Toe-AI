package game

import (
	"github.com/anilmisirlioglu/Tic-Tac-Toe-AI/math"
	"testing"
)

func TestNewBoardWriterBoardWriter(t *testing.T) {
	matrix := math.NewMatrix(3, 3, nil)

	board := NewBoardWriter(&matrix)

	value := board.matrix.GetElement(3, 3)
	if value == -1 {
		t.Errorf("Element not found.")
	}

	if value != 0 {
		t.Errorf("Actual: %d, Expected: 0", value)
	}
}
