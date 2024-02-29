package fit

type uint32Accumulator struct {
	accumuValue uint32
	lastValue   uint32
	mask        uint32
}

func uint32NewAccumulator(bits uint) *uint32Accumulator {
	return &uint32Accumulator{
		mask: (1 << bits) - 1,
	}
}

func (a *uint32Accumulator) accumulate(value uint32) uint32 {
	a.accumuValue += (value - a.lastValue) & a.mask
	a.lastValue = value
	return a.accumuValue
}
