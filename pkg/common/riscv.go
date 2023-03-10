// Copyright © 2015 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package common

import "fmt"

// immIFits reports whether immediate value x fits in nbits bits
// as a signed integer.
func ImmIFits[W int16 | int32 | int64](x W, nbits uint) bool {
	nbits--
	var min W = -1 << nbits
	var max W = 1<<nbits - 1
	return min <= x && x <= max
}

// immI extracts the signed integer of the specified size from an immediate.
func ImmI[W int16 | int32 | int64](imm W, nbits uint) (uint32, error) {
	if !ImmIFits(imm, nbits) {
		return 0, fmt.Errorf("signed immediate %d cannot fit in %d bits", imm, nbits)
	}
	return uint32(imm), nil
}

// signExtend sign extends val starting at bit bit.
func SignExtend[inp uint32 | int16 | int32 | int64, W uint32 | int16 | int32 | int64](val inp, bit uint) W {
	switch any(val).(type) {
	case int16:
		return W(val << (16 - bit) >> (16 - bit))
	case int32:
		return W(val << (32 - bit) >> (32 - bit))
	case uint32:
		return W(int32(val) << (32 - bit) >> (32 - bit))
	case int64:
		return W(val << (64 - bit) >> (64 - bit))
	default:
		// wot
		return 0
	}
}

// Split32BitImmediate splits a signed 32-bit immediate into a signed 20-bit
// upper immediate and a signed 12-bit lower immediate to be added to the upper
// result. For example, high may be used in LUI and low in a following ADDI to
// generate a full 32-bit constant.
func Split32BitImmediate[inp int64, outp int32](imm inp) (low, high outp, err error) {
	if !ImmIFits(imm, 32) {
		return 0, 0, fmt.Errorf("immediate does not fit in 32 bits: %d", imm)
	}

	// Nothing special needs to be done if the immediate fits in 12 bits.
	if ImmIFits(imm, 12) {
		return outp(imm), 0, nil
	}

	high = outp(imm) >> 12

	// The bottom 12 bits will be treated as signed.
	//
	// If that will result in a negative 12 bit number, add 1 to
	// our upper bits to adjust for the borrow.
	//
	// It is not possible for this increment to overflow. To
	// overflow, the 20 top bits would be 1, and the sign bit for
	// the low 12 bits would be set, in which case the entire 32
	// bit pattern fits in a 12 bit signed value.
	if imm&(1<<11) != 0 {
		high++
	}

	low = SignExtend[inp, outp](imm, 12)
	high = SignExtend[outp, outp](high, 20)

	return low, high, nil
}
