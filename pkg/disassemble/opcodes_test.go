package disassemble

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var opcodetests = []TableEntry{
	Entry("0x01d72223", 0x01d72223, SW(4, 14, 29, nil)),
	Entry("0xfc3f2223", 0xfc3f2223, SW(-60, 30, 3, nil)),
	Entry("slli", 0x01f51513, SLLI(10, 10, 31, nil)),
	Entry("beq", 0xfee786e3, BEQ(-20, 15, 14, nil)),
	Entry("addi", 0x67838393, ADDI(7, 7, 1656, nil)),
}

var _ = Describe("Opcodes", func() {

	var _ = DescribeTable("things", func(input int, expected Opcode) {
		out, err := Match(uint32(input))
		Expect(err).To(BeNil())

		Expect(out).To(Equal(expected))
	}, opcodetests)

})

func TestOpcodes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opcodes Suite")
}
