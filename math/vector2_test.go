package math

import "testing"

func TestNewVector2(t *testing.T) {
	v := NewVector2(4, 2)
	if v.X != 0 {
		t.Errorf("Actual %d, Excepted: 0", v.X)
	}
}

func TestVector2_String(t *testing.T) {
	v := NewVector2(2, 1)
	str := v.String()
	if str != "c2" {
		t.Errorf("Actual %s, Excepted: a2", str)
	}
}
