package test

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
	"github.com/hawkowl/biscuit/pkg/assemble"
	"github.com/hawkowl/biscuit/pkg/disassemble"
)

func TestRoundtrip(t *testing.T) {
	original := []uint32{
		0x01d72223,
		0x01c72423,
		0x00672623,
		0x01078793,
		0x01070713,
		0xfd079ce3,
		0x10089073,
		0x00269793,
		0x00f606b3,
		0x00b6a023,
		0x12050073, // fence.vma
		0x0000100f, // fence.i
		0x00c12083,
		0x01010113,
		0x00008067,
		0x0407f713,
		0x02070a63,
		0x0807f713,
		0x04071a63,
		0x00f00713,
		0x0807e793,
		0x04e59463,
		0x00269693,
		0x00d606b3,
		0x00f6a023,
		0x12050073, // fence.vma
		0x00c12083,
		0x01010113,
	}
	decoded := []disassemble.Opcode{}
	recoded := []uint32{}

	for _, i := range original {
		d, err := disassemble.Match(i)
		if err != nil {
			t.Fatal(err)
		}
		decoded = append(decoded, d)
	}

	// Now re-encode
	for _, i := range decoded {
		e, err := assemble.Encode(i.Opcode())
		if err != nil {
			t.Fatal(err)
		}

		recoded = append(recoded, e)
	}

	for _, i := range deep.Equal(original, recoded) {
		t.Error(i)
	}

	for i, f := range original {
		fmt.Printf("0x%032b - 0x%032b\n", f, recoded[i])
	}
}
