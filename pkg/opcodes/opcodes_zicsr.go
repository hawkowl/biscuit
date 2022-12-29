// AUTOGENERATED: DO NOT EDIT!
// Built from RISC-V ISA opcode definitions: https://github.com/riscv/riscv-opcodes

package opcodes

type OP_CSRRW struct {
	RD  uint32
	RS1 uint32
	CSR uint32
}

type OP_CSRRS struct {
	RD  uint32
	RS1 uint32
	CSR uint32
}

type OP_CSRRC struct {
	RD  uint32
	RS1 uint32
	CSR uint32
}

type OP_CSRRWI struct {
	RD   uint32
	CSR  uint32
	ZIMM uint32
}

type OP_CSRRSI struct {
	RD   uint32
	CSR  uint32
	ZIMM uint32
}

type OP_CSRRCI struct {
	RD   uint32
	CSR  uint32
	ZIMM uint32
}