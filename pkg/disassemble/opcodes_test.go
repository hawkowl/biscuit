package disassemble

import (
	"fmt"
	"testing"
)

func TestADDI(t *testing.T) {

	testCases := []struct {
		input    uint32
		expected OP_ADDI
	}{
		{
			input:    0x67838393,
			expected: ADDI(7, 7, 1656, nil),
		},
	}
	for _, tC := range testCases {
		funcName := fmt.Sprintf("%#v", tC.input)

		t.Run(funcName, func(t *testing.T) {
			out := DecodeADDI(tC.input)
			if out != tC.expected {
				t.Errorf("got %v, expected %v", out, tC.expected)
			}
		})
	}
}

func TestBEQ(t *testing.T) {

	testCases := []struct {
		input    uint32
		expected OP_BEQ
	}{
		{
			input:    0xfee786e3,
			expected: BEQ(-20, 15, 14, nil),
		},
	}
	for _, tC := range testCases {
		funcName := fmt.Sprintf("%#v", tC.input)

		t.Run(funcName, func(t *testing.T) {
			out := DecodeBEQ(tC.input)
			if out != tC.expected {
				t.Errorf("got %v, expected %v", out, tC.expected)
			}
		})
	}
}
