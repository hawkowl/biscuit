package debuginfo

type AssertionVariety int

const (
	NoAssertion AssertionVariety = iota
	RegisterAssertion
	MemoryAssertion
	CSRAssertion
)

type assertion struct {
	Target        uint32
	Variety       AssertionVariety
	ExpectedValue uint32
	Tag           string
}

func (d assertion) AssertionType() AssertionVariety {
	return d.Variety
}

func (d assertion) AssertionTarget() uint32 {
	return d.Target
}

func (d assertion) AssertWord(value uint32) (bool, uint32, string) {
	if value != d.ExpectedValue {
		return false, d.ExpectedValue, d.Tag
	}
	return true, 0, d.Tag
}

func (d assertion) AssertHalfWord(value uint16) (bool, uint16, string) {
	if value != uint16(d.ExpectedValue) {
		return false, uint16(d.ExpectedValue), d.Tag
	}
	return true, 0, d.Tag
}

func Assertion(variety AssertionVariety, target uint32, expectedValue uint32, tag string) DebugInfo {
	return assertion{
		Variety:       variety,
		Target:        target,
		ExpectedValue: expectedValue,
		Tag:           tag,
	}
}
