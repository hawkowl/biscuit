package disassemble

import (
	"github.com/hawkowl/biscuit/pkg/common"
	"github.com/hawkowl/biscuit/pkg/debuginfo"
	"github.com/hawkowl/biscuit/pkg/opcodes"
)

func field(in, start, length uint32) uint32 {
	return (in >> start) << (32 - length) >> (32 - length)
}

type Opcode interface {
	Describe() string
	Opcode() opcodes.Opcode
	DebugInfo() debuginfo.DebugInfo
}

func DEC_BIMM12(inst uint32) int32 {
	return common.SignExtend[uint32, int32](
		field(inst, 7, 1)<<11|field(inst, 8, 4)<<1|field(inst, 25, 6)<<5|field(inst, 31, 1)<<12,
		13,
	)
}

func DEC_IMM12HILO(inst uint32) int32 {
	return common.SignExtend[uint32, int32](
		field(inst, 25, 7)<<4|field(inst, 7, 4),
		12,
	)
}

func DEC_IMM12(inst uint32) int32 {
	return common.SignExtend[uint32, int32](
		field(inst, 20, 12),
		12,
	)
}

func DEC_IMM20(inst uint32) int32 {
	return common.SignExtend[uint32, int32](
		field(inst, 12, 20),
		20,
	)
}

func DEC_JIMM20(inst uint32) int32 {
	return common.SignExtend[uint32, int32](
		field(inst, 31, 1)<<20|field(inst, 21, 10)<<1|field(inst, 19, 1)<<11|field(inst, 12, 8)<<12,
		21,
	)
}

func DEC_FM(inst uint32) uint32 {
	return field(inst, 28, 4)
}

func DEC_PRED(inst uint32) uint32 {
	return field(inst, 24, 4)
}

func DEC_SUCC(inst uint32) uint32 {
	return field(inst, 20, 4)
}

func DEC_SHAMTW(inst uint32) uint32 {
	return field(inst, 20, 5)
}

func DEC_CSR(inst uint32) uint32 {
	return field(inst, 20, 12)
}

func DEC_ZIMM(inst uint32) uint32 {
	return field(inst, 15, 5)
}

func DEC_RD(inst uint32) uint32 {
	return field(inst, 7, 5)
}

func DEC_RS1(inst uint32) uint32 {
	return field(inst, 15, 5)
}

func DEC_RS2(inst uint32) uint32 {
	return field(inst, 20, 5)
}

type OP_ILLEGAL struct {
	opcodes.OP_ILLEGAL
}

func (o OP_ILLEGAL) Describe() string {
	return "ILLEGAL"
}

func (o OP_ILLEGAL) Opcode() opcodes.Opcode {
	return o.OP_ILLEGAL
}

func (o OP_ILLEGAL) DebugInfo() debuginfo.DebugInfo {
	return nil
}

func ILLEGAL() OP_ILLEGAL {
	return OP_ILLEGAL{
		opcodes.OP_ILLEGAL{},
	}
}

func NOP(debuginfo debuginfo.DebugInfo) Opcode {
	return ADDI(0, 0, 0, debuginfo)
}
