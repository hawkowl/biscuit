// AUTOGENERATED: DO NOT EDIT!
// Built from RISC-V ISA opcode definitions: https://github.com/riscv/riscv-opcodes

package opcodes

const (
	MASK_LUI     uint32 = 0x0000007F
	MATCH_LUI    uint32 = 0x00000037
	MASK_AUIPC   uint32 = 0x0000007F
	MATCH_AUIPC  uint32 = 0x00000017
	MASK_JAL     uint32 = 0x0000007F
	MATCH_JAL    uint32 = 0x0000006F
	MASK_JALR    uint32 = 0x0000707F
	MATCH_JALR   uint32 = 0x00000067
	MASK_BEQ     uint32 = 0x0000707F
	MATCH_BEQ    uint32 = 0x00000063
	MASK_BNE     uint32 = 0x0000707F
	MATCH_BNE    uint32 = 0x00001063
	MASK_BLT     uint32 = 0x0000707F
	MATCH_BLT    uint32 = 0x00004063
	MASK_BGE     uint32 = 0x0000707F
	MATCH_BGE    uint32 = 0x00005063
	MASK_BLTU    uint32 = 0x0000707F
	MATCH_BLTU   uint32 = 0x00006063
	MASK_BGEU    uint32 = 0x0000707F
	MATCH_BGEU   uint32 = 0x00007063
	MASK_LB      uint32 = 0x0000707F
	MATCH_LB     uint32 = 0x00000003
	MASK_LH      uint32 = 0x0000707F
	MATCH_LH     uint32 = 0x00001003
	MASK_LW      uint32 = 0x0000707F
	MATCH_LW     uint32 = 0x00002003
	MASK_LBU     uint32 = 0x0000707F
	MATCH_LBU    uint32 = 0x00004003
	MASK_LHU     uint32 = 0x0000707F
	MATCH_LHU    uint32 = 0x00005003
	MASK_SB      uint32 = 0x0000707F
	MATCH_SB     uint32 = 0x00000023
	MASK_SH      uint32 = 0x0000707F
	MATCH_SH     uint32 = 0x00001023
	MASK_SW      uint32 = 0x0000707F
	MATCH_SW     uint32 = 0x00002023
	MASK_ADDI    uint32 = 0x0000707F
	MATCH_ADDI   uint32 = 0x00000013
	MASK_SLTI    uint32 = 0x0000707F
	MATCH_SLTI   uint32 = 0x00002013
	MASK_SLTIU   uint32 = 0x0000707F
	MATCH_SLTIU  uint32 = 0x00003013
	MASK_XORI    uint32 = 0x0000707F
	MATCH_XORI   uint32 = 0x00004013
	MASK_ORI     uint32 = 0x0000707F
	MATCH_ORI    uint32 = 0x00006013
	MASK_ANDI    uint32 = 0x0000707F
	MATCH_ANDI   uint32 = 0x00007013
	MASK_ADD     uint32 = 0xFE00707F
	MATCH_ADD    uint32 = 0x00000033
	MASK_SUB     uint32 = 0xFE00707F
	MATCH_SUB    uint32 = 0x40000033
	MASK_SLL     uint32 = 0xFE00707F
	MATCH_SLL    uint32 = 0x00001033
	MASK_SLT     uint32 = 0xFE00707F
	MATCH_SLT    uint32 = 0x00002033
	MASK_SLTU    uint32 = 0xFE00707F
	MATCH_SLTU   uint32 = 0x00003033
	MASK_XOR     uint32 = 0xFE00707F
	MATCH_XOR    uint32 = 0x00004033
	MASK_SRL     uint32 = 0xFE00707F
	MATCH_SRL    uint32 = 0x00005033
	MASK_SRA     uint32 = 0xFE00707F
	MATCH_SRA    uint32 = 0x40005033
	MASK_OR      uint32 = 0xFE00707F
	MATCH_OR     uint32 = 0x00006033
	MASK_AND     uint32 = 0xFE00707F
	MATCH_AND    uint32 = 0x00007033
	MASK_FENCE   uint32 = 0x0000707F
	MATCH_FENCE  uint32 = 0x0000000F
	MASK_ECALL   uint32 = 0xFFFFFFFF
	MATCH_ECALL  uint32 = 0x00000073
	MASK_EBREAK  uint32 = 0xFFFFFFFF
	MATCH_EBREAK uint32 = 0x00100073
	MASK_SLLI    uint32 = 0xFE00707F
	MATCH_SLLI   uint32 = 0x00001013
	MASK_SRLI    uint32 = 0xFE00707F
	MATCH_SRLI   uint32 = 0x00005013
	MASK_SRAI    uint32 = 0xFE00707F
	MATCH_SRAI   uint32 = 0x40005013
)

type OP_LUI struct {
	RD    uint32
	IMM20 int32
}

func (o OP_LUI) Components() Components {
	return Components{
		RD:    o.RD,
		IMM20: o.IMM20,
	}
}

type OP_AUIPC struct {
	RD    uint32
	IMM20 int32
}

func (o OP_AUIPC) Components() Components {
	return Components{
		RD:    o.RD,
		IMM20: o.IMM20,
	}
}

type OP_JAL struct {
	RD     uint32
	JIMM20 int32
}

func (o OP_JAL) Components() Components {
	return Components{
		RD:     o.RD,
		JIMM20: o.JIMM20,
	}
}

type OP_JALR struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_JALR) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_BEQ struct {
	BIMM12 int32
	RS1    uint32
	RS2    uint32
}

func (o OP_BEQ) Components() Components {
	return Components{
		BIMM12: o.BIMM12,
		RS1:    o.RS1,
		RS2:    o.RS2,
	}
}

type OP_BNE struct {
	BIMM12 int32
	RS1    uint32
	RS2    uint32
}

func (o OP_BNE) Components() Components {
	return Components{
		BIMM12: o.BIMM12,
		RS1:    o.RS1,
		RS2:    o.RS2,
	}
}

type OP_BLT struct {
	BIMM12 int32
	RS1    uint32
	RS2    uint32
}

func (o OP_BLT) Components() Components {
	return Components{
		BIMM12: o.BIMM12,
		RS1:    o.RS1,
		RS2:    o.RS2,
	}
}

type OP_BGE struct {
	BIMM12 int32
	RS1    uint32
	RS2    uint32
}

func (o OP_BGE) Components() Components {
	return Components{
		BIMM12: o.BIMM12,
		RS1:    o.RS1,
		RS2:    o.RS2,
	}
}

type OP_BLTU struct {
	BIMM12 int32
	RS1    uint32
	RS2    uint32
}

func (o OP_BLTU) Components() Components {
	return Components{
		BIMM12: o.BIMM12,
		RS1:    o.RS1,
		RS2:    o.RS2,
	}
}

type OP_BGEU struct {
	BIMM12 int32
	RS1    uint32
	RS2    uint32
}

func (o OP_BGEU) Components() Components {
	return Components{
		BIMM12: o.BIMM12,
		RS1:    o.RS1,
		RS2:    o.RS2,
	}
}

type OP_LB struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_LB) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_LH struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_LH) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_LW struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_LW) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_LBU struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_LBU) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_LHU struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_LHU) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_SB struct {
	IMM12 int32
	RS1   uint32
	RS2   uint32
}

func (o OP_SB) Components() Components {
	return Components{
		IMM12: o.IMM12,
		RS1:   o.RS1,
		RS2:   o.RS2,
	}
}

type OP_SH struct {
	IMM12 int32
	RS1   uint32
	RS2   uint32
}

func (o OP_SH) Components() Components {
	return Components{
		IMM12: o.IMM12,
		RS1:   o.RS1,
		RS2:   o.RS2,
	}
}

type OP_SW struct {
	IMM12 int32
	RS1   uint32
	RS2   uint32
}

func (o OP_SW) Components() Components {
	return Components{
		IMM12: o.IMM12,
		RS1:   o.RS1,
		RS2:   o.RS2,
	}
}

type OP_ADDI struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_ADDI) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_SLTI struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_SLTI) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_SLTIU struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_SLTIU) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_XORI struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_XORI) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_ORI struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_ORI) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_ANDI struct {
	RD    uint32
	RS1   uint32
	IMM12 int32
}

func (o OP_ANDI) Components() Components {
	return Components{
		RD:    o.RD,
		RS1:   o.RS1,
		IMM12: o.IMM12,
	}
}

type OP_ADD struct {
	RD  uint32
	RS1 uint32
	RS2 uint32
}

func (o OP_ADD) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		RS2: o.RS2,
	}
}

type OP_SUB struct {
	RD  uint32
	RS1 uint32
	RS2 uint32
}

func (o OP_SUB) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		RS2: o.RS2,
	}
}

type OP_SLL struct {
	RD  uint32
	RS1 uint32
	RS2 uint32
}

func (o OP_SLL) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		RS2: o.RS2,
	}
}

type OP_SLT struct {
	RD  uint32
	RS1 uint32
	RS2 uint32
}

func (o OP_SLT) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		RS2: o.RS2,
	}
}

type OP_SLTU struct {
	RD  uint32
	RS1 uint32
	RS2 uint32
}

func (o OP_SLTU) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		RS2: o.RS2,
	}
}

type OP_XOR struct {
	RD  uint32
	RS1 uint32
	RS2 uint32
}

func (o OP_XOR) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		RS2: o.RS2,
	}
}

type OP_SRL struct {
	RD  uint32
	RS1 uint32
	RS2 uint32
}

func (o OP_SRL) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		RS2: o.RS2,
	}
}

type OP_SRA struct {
	RD  uint32
	RS1 uint32
	RS2 uint32
}

func (o OP_SRA) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		RS2: o.RS2,
	}
}

type OP_OR struct {
	RD  uint32
	RS1 uint32
	RS2 uint32
}

func (o OP_OR) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		RS2: o.RS2,
	}
}

type OP_AND struct {
	RD  uint32
	RS1 uint32
	RS2 uint32
}

func (o OP_AND) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		RS2: o.RS2,
	}
}

type OP_FENCE struct {
	FM   uint32
	PRED uint32
	SUCC uint32
	RS1  uint32
	RD   uint32
}

func (o OP_FENCE) Components() Components {
	return Components{
		FM:   o.FM,
		PRED: o.PRED,
		SUCC: o.SUCC,
		RS1:  o.RS1,
		RD:   o.RD,
	}
}

type OP_ECALL struct {
}

func (o OP_ECALL) Components() Components {
	return Components{}
}

type OP_EBREAK struct {
}

func (o OP_EBREAK) Components() Components {
	return Components{}
}

type OP_SLLI struct {
	RD     uint32
	RS1    uint32
	SHAMTW uint32
}

func (o OP_SLLI) Components() Components {
	return Components{
		RD:     o.RD,
		RS1:    o.RS1,
		SHAMTW: o.SHAMTW,
	}
}

type OP_SRLI struct {
	RD     uint32
	RS1    uint32
	SHAMTW uint32
}

func (o OP_SRLI) Components() Components {
	return Components{
		RD:     o.RD,
		RS1:    o.RS1,
		SHAMTW: o.SHAMTW,
	}
}

type OP_SRAI struct {
	RD     uint32
	RS1    uint32
	SHAMTW uint32
}

func (o OP_SRAI) Components() Components {
	return Components{
		RD:     o.RD,
		RS1:    o.RS1,
		SHAMTW: o.SHAMTW,
	}
}
