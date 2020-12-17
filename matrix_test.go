package main

import "testing"

func TestNewMatrix(t *testing.T) {
	m := NewMatrix(3, 3, MatrixMap{})
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Columns; c++ {
			item, found := m.Matrix[r][c]
			if !found {
				t.Fatalf("Actual: %d, Expected: 0", item)
			}
		}
	}
}

func TestMatrix_GetElement(t *testing.T) {
	m := NewMatrix(3, 3, MatrixMap{})
	item, found := m.GetElement(1, 2)
	if !found {
		t.Fatalf("Actual: %d, Expected: 0", item)
	}
}

func TestMatrix_SetElement(t *testing.T) {
	m := NewMatrix(3, 3, MatrixMap{})
	m.SetElement(1, 1, 1)

	item, found := m.GetElement(1, 2)
	if !found {
		t.Fatalf("Actual %d, Expected: 1", item)
	}
}
