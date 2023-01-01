package debuginfo

type DebugInfo interface {
	AssertionType() AssertionVariety
	AssertionTarget() uint32
	AssertWord(uint32) (ok bool, expected uint32, tag string)
	AssertHalfWord(uint16) (ok bool, expected uint16, tag string)
}
