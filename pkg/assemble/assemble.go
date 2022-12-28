package assemble

import (
	"fmt"
	"math/bits"

	"github.com/hawkowl/biscuit/pkg/common"
)

type Opcode interface {
	Encode() (uint32, error)
}

func FitsIn(inp uint32, nbits int) (uint32, error) {
	if bits.Len32(inp) > nbits {
		return 0, fmt.Errorf("%d does not fit in %d bits", inp, nbits)
	}

	return uint32(inp), nil
}

func GEN_IMM20(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 20)
	return i << 12, err
}

func GEN_IMM12(inp int32) (uint32, error) {
	i, err := common.ImmI(inp, 12)
	return i << 20, err
}

func GEN_IMM12HI(inp int32) (uint32, error) {
	i, err := common.ImmI(inp, 12)
	return (i >> 5) << 25, err
}

func GEN_IMM12LO(inp int32) (uint32, error) {
	i, err := common.ImmI(inp, 12)
	return (i & 0x1f) << 7, err
}

func GEN_BIMM12HI(inp int32) (uint32, error) {
	i, err := common.ImmI(inp, 13)
	return (i>>12)<<31 | ((i>>5)&0x3f)<<25, err
}

func GEN_BIMM12LO(inp int32) (uint32, error) {
	i, err := common.ImmI(inp, 13)
	return ((i>>1)&0xf)<<8 | ((i>>11)&0x1)<<7, err
}

func GEN_JIMM20(inp int32) (uint32, error) {
	i, err := common.ImmI(inp, 21)
	return (i>>20)<<31 | ((i>>1)&0x3ff)<<21 | ((i>>11)&0x1)<<20 | ((i>>12)&0xff)<<12, err
}

func GEN_SUCC(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return (i & 0xF) << 20, err
}

func GEN_PRED(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return ((i >> 4) & 0xF) << 24, err
}

func GEN_FM(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return ((i >> 8) & 0xF) << 28, err
}

func GEN_RD(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 5)
	return i << 7, err
}

func GEN_RS1(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 5)
	return i << 15, err
}

func GEN_RS2(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 5)
	return i << 20, err
}

func GEN_CSR(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return i << 20, err
}

func GEN_ZIMM(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 5)
	return i << 20, err
}

func GEN_SHAMTW(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 5)
	return i << 20, err
}
