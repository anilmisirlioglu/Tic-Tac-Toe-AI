package math

type MatrixMap = map[int]map[int]int

type IMatrix interface {
	Set(array MatrixMap)
	SetElement(row, column, value int) bool
	GetElement(row, column int) (int, bool)
}

type Matrix struct {
	Rows    int
	Columns int
	Map     MatrixMap
}

func (m Matrix) Set(array MatrixMap) {
	for r := 0; r < m.Rows; r++ {
		m.Map[r] = map[int]int{}
		for c := 0; c < m.Columns; c++ {
			item := 0
			if value, found := array[r][c]; found {
				item = value
			}

			m.Map[r][c] = item
		}
	}
}

func (m Matrix) SetElement(row, column, value int) bool {
	if row > m.Rows || row < 0 || column > m.Columns || column < 0 {
		return false
	}

	m.Map[row][column] = value
	return true
}

func (m Matrix) GetElement(row, column int) int {
	if row > m.Rows || row < 0 || column > m.Columns || column < 0 {
		return -1
	}

	return m.Map[row][column]
}

func NewMatrix(rows, columns int, array MatrixMap) Matrix {
	if rows <= 0 {
		rows = 1
	}

	if columns <= 0 {
		columns = 1
	}

	matrix := Matrix{
		Rows:    rows,
		Columns: columns,
		Map:     make(MatrixMap),
	}

	matrix.Set(array)

	return matrix
}
