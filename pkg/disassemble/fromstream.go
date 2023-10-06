package disassemble

import (
	"fmt"
	"io"
)

func Decode(buf io.Reader) ([]Opcode, error) {
	i := make([]Opcode, 0)
	b := make([]byte, 1)
	location := 0

	for {
		r := uint32(0)

		for _i := 0; _i < 4; _i++ {
			n, err := buf.Read(b)
			if n != 1 {
				return i, fmt.Errorf("what? got %d", n)
			}
			if err != nil {
				if err == io.EOF {
					//fmt.Printf("EOL")
					return i, nil
				}
				return nil, err
			}
			location += 1

			r = r >> 8
			r = r + (uint32(b[0]) << 24)

		}

		dis, err := Match(r)
		if err != nil {
			// zeroes?
			if r == 0 {
				return i, nil
			}
			return nil, fmt.Errorf("at offset %d %x: %w", location, location, err)
		}

		i = append(i, dis)
	}
}
