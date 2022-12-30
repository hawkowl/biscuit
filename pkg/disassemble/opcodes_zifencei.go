// AUTOGENERATED: DO NOT EDIT!
// Built from RISC-V ISA opcode definitions: https://github.com/riscv/riscv-opcodes

package disassemble

import "github.com/hawkowl/biscuit/pkg/opcodes"

type OP_FENCE_I struct {
	opcodes.OP_FENCE_I
}

func (o OP_FENCE_I) Describe() string {
	return "FENCE_I"
}

func (o OP_FENCE_I) Opcode() opcodes.Opcode {
	return o.OP_FENCE_I
}

func FENCE_I(IMM12 int32, RS1 uint32, RD uint32, debug DebugInfo) OP_FENCE_I {
	return OP_FENCE_I{
		opcodes.OP_FENCE_I{
			IMM12: IMM12,
			RS1:   RS1,
			RD:    RD,
		},
	}
}

func DecodeFENCE_I(inst uint32) OP_FENCE_I {
	r := opcodes.OP_FENCE_I{}

	r.IMM12 = DEC_IMM12(inst)
	r.RS1 = DEC_RS1(inst)
	r.RD = DEC_RD(inst)
	return OP_FENCE_I{
		OP_FENCE_I: r,
	}
}
