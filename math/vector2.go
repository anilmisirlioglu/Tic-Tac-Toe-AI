package math

import "strconv"

type Vector2 struct {
	X, Y int
}

func (v Vector2) String() string {
	return []string{"a", "b", "c"}[v.X] + strconv.Itoa(v.Y+1)
}

func NewVector2(x, y int) Vector2 {
	checkAxisRange(&x)
	checkAxisRange(&y)

	return Vector2{
		X: x,
		Y: y,
	}
}

func checkAxisRange(axis *int) {
	if *axis > 2 || *axis < 0 {
		*axis = 0
	}
}
