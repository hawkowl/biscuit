package main

import (
	"debug/elf"
	"fmt"
	"os"

	"github.com/hawkowl/biscuit/pkg/disassemble"
	"github.com/hawkowl/biscuit/pkg/elfsupport"
)

func openELF(path string) (elfsupport.ELF, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return elfsupport.Open(f)
}

func formatSection(el elfsupport.ELF, p *elf.Prog) ([]string, error) {
	opcodes, err := disassemble.Decode(p.Open())
	if err != nil {
		return nil, err
	}

	r := make([]string, 0, len(opcodes))

	for i, o := range opcodes {
		r = append(r, fmt.Sprintf("%#x: %s", p.Vaddr+uint64(i*4), o.String()))
	}

	return r, nil
}

func main() {
	path := os.Args[1]

	e, err := openELF(path)
	if err != nil {
		panic(err)
	}

	var p *elf.Prog

	for _, prog := range e.Progs() {
		if e.Entry() == prog.Vaddr {
			p = prog
		}
	}

	r, err := formatSection(e, p)
	if err != nil {
		panic(err)
	}

	for _, i := range r {
		fmt.Println(i)
	}

	return
}
