// AUTOGENERATED: DO NOT EDIT!
// Built from RISC-V ISA opcode definitions: https://github.com/riscv/riscv-opcodes

package assemble

import (
	"fmt"
	"github.com/hawkowl/biscuit/pkg/opcodes"
)

func Encode(inp opcodes.Opcode) (uint32, error) {
	switch r := inp.(type) {
	case opcodes.OP_LUI:
		return EncodeLUI(r)
	case opcodes.OP_AUIPC:
		return EncodeAUIPC(r)
	case opcodes.OP_JAL:
		return EncodeJAL(r)
	case opcodes.OP_JALR:
		return EncodeJALR(r)
	case opcodes.OP_BEQ:
		return EncodeBEQ(r)
	case opcodes.OP_BNE:
		return EncodeBNE(r)
	case opcodes.OP_BLT:
		return EncodeBLT(r)
	case opcodes.OP_BGE:
		return EncodeBGE(r)
	case opcodes.OP_BLTU:
		return EncodeBLTU(r)
	case opcodes.OP_BGEU:
		return EncodeBGEU(r)
	case opcodes.OP_LB:
		return EncodeLB(r)
	case opcodes.OP_LH:
		return EncodeLH(r)
	case opcodes.OP_LW:
		return EncodeLW(r)
	case opcodes.OP_LBU:
		return EncodeLBU(r)
	case opcodes.OP_LHU:
		return EncodeLHU(r)
	case opcodes.OP_SB:
		return EncodeSB(r)
	case opcodes.OP_SH:
		return EncodeSH(r)
	case opcodes.OP_SW:
		return EncodeSW(r)
	case opcodes.OP_ADDI:
		return EncodeADDI(r)
	case opcodes.OP_SLTI:
		return EncodeSLTI(r)
	case opcodes.OP_SLTIU:
		return EncodeSLTIU(r)
	case opcodes.OP_XORI:
		return EncodeXORI(r)
	case opcodes.OP_ORI:
		return EncodeORI(r)
	case opcodes.OP_ANDI:
		return EncodeANDI(r)
	case opcodes.OP_ADD:
		return EncodeADD(r)
	case opcodes.OP_SUB:
		return EncodeSUB(r)
	case opcodes.OP_SLL:
		return EncodeSLL(r)
	case opcodes.OP_SLT:
		return EncodeSLT(r)
	case opcodes.OP_SLTU:
		return EncodeSLTU(r)
	case opcodes.OP_XOR:
		return EncodeXOR(r)
	case opcodes.OP_SRL:
		return EncodeSRL(r)
	case opcodes.OP_SRA:
		return EncodeSRA(r)
	case opcodes.OP_OR:
		return EncodeOR(r)
	case opcodes.OP_AND:
		return EncodeAND(r)
	case opcodes.OP_FENCE:
		return EncodeFENCE(r)
	case opcodes.OP_ECALL:
		return EncodeECALL(r)
	case opcodes.OP_EBREAK:
		return EncodeEBREAK(r)
	case opcodes.OP_SLLI:
		return EncodeSLLI(r)
	case opcodes.OP_SRLI:
		return EncodeSRLI(r)
	case opcodes.OP_SRAI:
		return EncodeSRAI(r)
	case opcodes.OP_SFENCE_VMA:
		return EncodeSFENCE_VMA(r)
	case opcodes.OP_SRET:
		return EncodeSRET(r)
	case opcodes.OP_CSRRW:
		return EncodeCSRRW(r)
	case opcodes.OP_CSRRS:
		return EncodeCSRRS(r)
	case opcodes.OP_CSRRC:
		return EncodeCSRRC(r)
	case opcodes.OP_CSRRWI:
		return EncodeCSRRWI(r)
	case opcodes.OP_CSRRSI:
		return EncodeCSRRSI(r)
	case opcodes.OP_CSRRCI:
		return EncodeCSRRCI(r)
	case opcodes.OP_FENCE_I:
		return EncodeFENCE_I(r)
	case opcodes.OP_ILLEGAL:
		return EncodeILLEGAL(r)
	default:
		return 0, fmt.Errorf("unknown %v", inp)
	}
}