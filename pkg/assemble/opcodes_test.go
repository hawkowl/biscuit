package assemble

import (
	"reflect"
	"testing"
)

func TestOpcodes(t *testing.T) {
	o := OP_beq{
		IMM12: -20,
		RS2:   14,
		RS1:   15,
	}

	r, err := o.Encode()

	if err != nil {
		t.Fatal(err)
	}

	expected := uint32(0xfee786e3)

	if r != expected {
		t.Fatalf("%032b not %032b", r, expected)
	}
}

func TestOpcodePortions(t *testing.T) {
	testCases := []struct {
		input  uint32
		f      func(uint32) (uint32, error)
		result uint32
	}{
		{
			input:  0b11110000,
			f:      GEN_pred,
			result: 0xF000000,
		},
		{
			input:  0b1111,
			f:      GEN_succ,
			result: 0xF00000,
		},
		{
			input:  0b111100000000,
			f:      GEN_fm,
			result: 0xF0000000,
		},
		{
			input:  0b11111,
			f:      GEN_rd,
			result: 0b111110000000,
		},
		{
			input:  0b11111,
			f:      GEN_rs1,
			result: 0b11111000000000000000,
		},
		{
			input:  0b11111,
			f:      GEN_rs2,
			result: 0b1111100000000000000000000,
		},
	}
	for _, tC := range testCases {

		funcName := reflect.ValueOf(tC.f).Kind().String()

		t.Run(funcName, func(t *testing.T) {
			out, err := tC.f(tC.input)
			if err != nil {
				t.Fatal(err)
			}
			if out != tC.result {
				t.Errorf("got %032b, wanted %032b", out, tC.result)
			}
		})
	}
}
