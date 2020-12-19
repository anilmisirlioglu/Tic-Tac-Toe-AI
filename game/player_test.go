package game

import "testing"

func Test_GetSymbolByInt(t *testing.T) {
	cpu := CPU{Symbol: X}
	human := Human{Symbol: O}

	cpuSymbol := cpu.GetSymbolByInt()
	if cpuSymbol != XIndex {
		t.Errorf("Actual %d, Expected: %d", cpuSymbol, XIndex)
	}

	humanSymbol := human.GetSymbolByInt()
	if humanSymbol != OIndex {
		t.Errorf("Actual %d, Expected: %d", humanSymbol, OIndex)
	}
}
