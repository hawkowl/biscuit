package assemble

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOpcodes(t *testing.T) {
	testCases := []struct {
		input    Opcode
		expected uint32
	}{
		{
			input:    BEQ(-20, 15, 14),
			expected: 0xfee786e3,
		},
		{
			input:    ADDI(7, 7, 1656),
			expected: 0x67838393,
		},
		{
			input:    SLLI(14, 10, 2),
			expected: 0x00251713,
		},
		{
			input:    JAL(1, -96),
			expected: 0xfa1ff0ef,
		},
		{
			input:    JAL(0, 9128),
			expected: 0x3a80206f,
		},
	}
	for _, tC := range testCases {
		funcName := fmt.Sprintf("%#v", tC.input)

		t.Run(funcName, func(t *testing.T) {
			out, err := tC.input.Encode()
			if err != nil {
				t.Fatal(err)
			}
			if out != tC.expected {
				t.Errorf("got %032b (0x%08X), wanted %032b (0x%08X)", out, out, tC.expected, tC.expected)
			}
		})
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
			f:      GEN_PRED,
			result: 0xF000000,
		},
		{
			input:  0b1111,
			f:      GEN_SUCC,
			result: 0xF00000,
		},
		{
			input:  0b111100000000,
			f:      GEN_FM,
			result: 0xF0000000,
		},
		{
			input:  0b11111,
			f:      GEN_RD,
			result: 0b111110000000,
		},
		{
			input:  0b11111,
			f:      GEN_RS1,
			result: 0b11111000000000000000,
		},
		{
			input:  0b11111,
			f:      GEN_RS2,
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
