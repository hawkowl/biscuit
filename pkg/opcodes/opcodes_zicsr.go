// AUTOGENERATED: DO NOT EDIT!
// Built from RISC-V ISA opcode definitions: https://github.com/riscv/riscv-opcodes

package opcodes

const (
	MASK_CSRRW   uint32 = 0x0000707F
	MATCH_CSRRW  uint32 = 0x00001073
	MASK_CSRRS   uint32 = 0x0000707F
	MATCH_CSRRS  uint32 = 0x00002073
	MASK_CSRRC   uint32 = 0x0000707F
	MATCH_CSRRC  uint32 = 0x00003073
	MASK_CSRRWI  uint32 = 0x0000707F
	MATCH_CSRRWI uint32 = 0x00005073
	MASK_CSRRSI  uint32 = 0x0000707F
	MATCH_CSRRSI uint32 = 0x00006073
	MASK_CSRRCI  uint32 = 0x0000707F
	MATCH_CSRRCI uint32 = 0x00007073
)

type OP_CSRRW struct {
	RD  uint32
	RS1 uint32
	CSR uint32
}

func (o OP_CSRRW) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		CSR: o.CSR,
	}
}

func (o OP_CSRRW) Instruction() string {
	return "CSRRW"
}

type OP_CSRRS struct {
	RD  uint32
	RS1 uint32
	CSR uint32
}

func (o OP_CSRRS) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		CSR: o.CSR,
	}
}

func (o OP_CSRRS) Instruction() string {
	return "CSRRS"
}

type OP_CSRRC struct {
	RD  uint32
	RS1 uint32
	CSR uint32
}

func (o OP_CSRRC) Components() Components {
	return Components{
		RD:  o.RD,
		RS1: o.RS1,
		CSR: o.CSR,
	}
}

func (o OP_CSRRC) Instruction() string {
	return "CSRRC"
}

type OP_CSRRWI struct {
	RD   uint32
	CSR  uint32
	ZIMM uint32
}

func (o OP_CSRRWI) Components() Components {
	return Components{
		RD:   o.RD,
		CSR:  o.CSR,
		ZIMM: o.ZIMM,
	}
}

func (o OP_CSRRWI) Instruction() string {
	return "CSRRWI"
}

type OP_CSRRSI struct {
	RD   uint32
	CSR  uint32
	ZIMM uint32
}

func (o OP_CSRRSI) Components() Components {
	return Components{
		RD:   o.RD,
		CSR:  o.CSR,
		ZIMM: o.ZIMM,
	}
}

func (o OP_CSRRSI) Instruction() string {
	return "CSRRSI"
}

type OP_CSRRCI struct {
	RD   uint32
	CSR  uint32
	ZIMM uint32
}

func (o OP_CSRRCI) Components() Components {
	return Components{
		RD:   o.RD,
		CSR:  o.CSR,
		ZIMM: o.ZIMM,
	}
}

func (o OP_CSRRCI) Instruction() string {
	return "CSRRCI"
}
