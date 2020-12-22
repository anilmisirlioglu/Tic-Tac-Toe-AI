package game

const (
	Prefix = "[*]"

	NextLineNone     = 0x1
	NextLineBefore   = 0x2
	NextLineBehind   = 0x4
	NextLineMultiply = NextLineBefore ^ NextLineBehind

	XWon       = 0x1
	OWon       = 0x2
	Draw       = 0x4
	Unfinished = 0x8

	XIndex    = 1
	OIndex    = 2
	NullIndex = 0

	X    = "X"
	O    = "O"
	NULL = "."

	Vertical   = 3
	Horizontal = 3
)

var SymbolIndexes = map[int]string{
	XIndex: X,
	OIndex: O,
}
