package assemble

import (
	"fmt"
	"math/bits"
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

func GEN_imm20(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 20)
	return i << 12, err
}

func GEN_imm12(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return i << 20, err
}

func GEN_imm12hi(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return (i >> 5) << 25, err
}

func GEN_imm12lo(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return (i & 0x1f) << 7, err
}

func GEN_bimm12hi(inp int32) (uint32, error) {
	i, err := immI(inp, 13)
	return (i>>12)<<31 | ((i>>5)&0x3f)<<25, err
}

func GEN_bimm12lo(inp int32) (uint32, error) {
	i, err := immI(inp, 13)
	return ((i>>1)&0xf)<<8 | ((i>>11)&0x1)<<7, err
}

func GEN_jimm20(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 20)
	return (i>>20)<<31 | ((i>>1)&0x3ff)<<21 | ((i>>11)&0x1)<<20 | ((i>>12)&0xff)<<12, err
}

func GEN_succ(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return (i & 0xF) << 20, err
}

func GEN_pred(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return ((i >> 4) & 0xF) << 24, err
}

func GEN_fm(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return ((i >> 8) & 0xF) << 28, err
}

func GEN_rd(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 5)
	return i << 7, err
}

func GEN_rs1(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 5)
	return i << 15, err
}

func GEN_rs2(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 5)
	return i << 20, err
}

func GEN_csr(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 12)
	return i << 20, err
}

func GEN_zimm(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 5)
	return i << 20, err
}

func GEN_shamtw(inp uint32) (uint32, error) {
	i, err := FitsIn(inp, 5)
	return i << 20, err
}
