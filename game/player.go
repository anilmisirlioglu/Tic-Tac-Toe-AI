package game

type Player interface {
	GetSymbolByInt() int
	GetName() string
}

type Human struct {
	Symbol string
}

func (h Human) GetName() string {
	return "Human"
}

func (h Human) GetSymbolByInt() int {
	if h.Symbol == X {
		return XIndex
	}

	return OIndex
}

type CPU struct {
	Symbol string
}

func (c CPU) GetName() string {
	return "Computer"
}

func (c CPU) GetSymbolByInt() int {
	if c.Symbol == X {
		return XIndex
	}

	return OIndex
}
