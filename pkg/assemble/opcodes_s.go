// AUTOGENERATED: DO NOT EDIT!
// Built from RISC-V ISA opcode definitions: https://github.com/riscv/riscv-opcodes

package assemble

import "github.com/hawkowl/biscuit/pkg/opcodes"

func EncodeSFENCE_VMA(o opcodes.OP_SFENCE_VMA) (uint32, error) {
	RS1, err := GEN_RS1(o.RS1)
	if err != nil {
		return 0, err
	}
	RS2, err := GEN_RS2(o.RS2)
	if err != nil {
		return 0, err
	}
	return 0x12000073 | RS1 | RS2, nil
}

func EncodeSRET(o opcodes.OP_SRET) (uint32, error) {
	return 0x10200073, nil
}
