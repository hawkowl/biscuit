package elfsupport

import (
	"debug/elf"
	"io"
)

type _elf struct {
	entry uint64
	progs []*elf.Prog

	syms []elf.Symbol
}

type ELF interface {
	Symbols() []elf.Symbol
	Progs() []*elf.Prog
	Entry() uint64
}

func Open(inp io.ReaderAt) (ELF, error) {
	e, err := elf.NewFile(inp)
	if err != nil {
		return nil, err
	}

	syms, err := e.Symbols()
	if err != nil {
		return nil, err
	}

	return &_elf{
		entry: e.Entry,
		progs: e.Progs,
		syms:  syms,
	}, nil
}

func (e *_elf) Symbols() []elf.Symbol {
	return e.syms
}

func (e *_elf) Progs() []*elf.Prog {
	return e.progs
}

func (e *_elf) Entry() uint64 {
	return e.entry
}
