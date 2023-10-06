package assemble

import (
	"reflect"
	"testing"

	"github.com/hawkowl/biscuit/pkg/opcodes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var opcodetests = []TableEntry{
	Entry("", 0xfee786e3, opcodes.OP_BEQ{RS1: 15, RS2: 14, BIMM12: -20}),
	Entry("", 0x67838393, opcodes.OP_ADDI{RD: 7, RS1: 7, IMM12: 1656}),
	Entry("", 0x00251713, opcodes.OP_SLLI{RD: 14, RS1: 10, SHAMTW: 2}),
	Entry("", 0xfa1ff0ef, opcodes.OP_JAL{RD: 1, JIMM20: -96}),
	Entry("", 0x3a80206f, opcodes.OP_JAL{RD: 0, JIMM20: 9128}),
}

var _ = Describe("Opcodes", func() {

	var _ = DescribeTable("things", func(expected int, op opcodes.Opcode) {

		out, err := Encode(op)
		Expect(err).To(BeNil())

		Expect(out).To(BeEquivalentTo(expected), "got %032b (0x%08X), wanted %032b (0x%08X)", out, out, expected, expected)
	}, opcodetests)

})

func TestOpcodes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opcodes Suite")
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
