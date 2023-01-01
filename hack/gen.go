package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"log"
	"math/bits"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

const DEFAULT_EXTENSIONS = "i,s,zicsr,zifencei,system"

type Opcode struct {
	Name   string
	Mask   uint32
	Match  uint32
	Args   [][3]string
	Fields [][2]string
}

type Data struct {
	Opcodes []Opcode
}

var msblsb, args *regexp.Regexp

func fileProcess(data *Data, filename string) error {
	d, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(d), "\n")

	for _, l := range lines {
		ln := strings.TrimSpace(l)
		ln, _, _ = strings.Cut(ln, "#")

		// don't use blanks or care about psuedo-ops rn
		if ln == "" || strings.HasPrefix(ln, "$import") {
			continue
		}

		if strings.HasPrefix(ln, "$pseudo_op") {
			lns := strings.SplitN(ln, " ", 3)
			oploc := strings.SplitN(lns[1], "::", 2)

			if strings.HasPrefix(oploc[0], "rv128") || strings.HasPrefix(oploc[0], "rv64") {
				log.Printf("processing psuedo-op %s", oploc[1])
				ln = lns[2]
			} else {
				continue
			}
		}

		op := Opcode{
			Mask: uint32(2 ^ 32),
		}
		fields := make(map[string]int)
		fieldNames := make([][3]string, 0)

		for i, f := range strings.Fields(ln) {
			if i == 0 {
				op.Name = strings.ToUpper(strings.Replace(f, ".", "_", -1))
				continue
			}

			subs := msblsb.FindStringSubmatch(f)
			if subs == nil {
				found := false
				for n, fieldname := range args.FindStringSubmatch(f) {
					if fieldname == "" || n == 0 {
						continue
					}

					fieldname = strings.ToUpper(fieldname)
					op.Fields = append(op.Fields, [2]string{strings.ToUpper(f), fieldname})

					_, ok := fields[fieldname]
					if !ok {
						fields[fieldname] = 1
						if fieldname == "BIMM12" || fieldname == "IMM12" {
							fieldNames = append(fieldNames, [3]string{fieldname, "int32", "12"})
						} else if fieldname == "JIMM20" || fieldname == "IMM20" {
							fieldNames = append(fieldNames, [3]string{fieldname, "int32", "20"})
						} else {
							fieldNames = append(fieldNames, [3]string{fieldname, "uint32", "12"})
						}
					}

					found = true
				}

				if !found {
					log.Printf("!! couldn't find field for %s\n", f)
				}
			} else {
				msb, err := strconv.Atoi(subs[1])
				if err != nil {
					return err
				}
				lsb, err := strconv.Atoi(subs[2])
				if err != nil {
					return err
				}
				value, err := strconv.ParseUint(subs[3], 0, msb-lsb+1)
				if err != nil {
					return err
				}

				op.Match = op.Match | bits.RotateLeft32(uint32(value), lsb)
				op.Mask = op.Mask | bits.RotateLeft32(uint32(1<<(msb-lsb+1)-1), lsb)
			}
		}
		op.Args = fieldNames
		data.Opcodes = append(data.Opcodes, op)

	}

	return nil
}

func writeFile(tmpl *template.Template, data *Data, dest string) error {
	var out bytes.Buffer
	err := tmpl.Execute(&out, data)
	if err != nil {
		return err
	}

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format %s: %w", dest, err)
	}

	return os.WriteFile(dest, formatted, 0o666)
}

func process(opcodesPath string, extensions []string) error {
	msblsb = regexp.MustCompile(`\s*(?P<msb>\d+.?)\.\.(?P<lsb>\d+.?)\s*=\s*(?P<val>\d[\w]*)[\s$]*`)
	args = regexp.MustCompile(`\s?(?:(?:((?:j|b|z)?imm(?:12|20)?)(?:hi|lo)?)+|(r(?:s\d|d))+|(fm|pred|succ|csr|shamtw)+)\s?`)

	tmplEncode := template.Must(template.ParseFiles("hack/templates/opcode_encode.tmpl"))
	tmplDecode := template.Must(template.ParseFiles("hack/templates/opcode_decode.tmpl"))
	tmplDefs := template.Must(template.ParseFiles("hack/templates/opcode_defs.tmpl"))
	tmplAssemble := template.Must(template.ParseFiles("hack/templates/opcode_assemble.tmpl"))
	tmplMatch := template.Must(template.ParseFiles("hack/templates/opcode_match.tmpl"))

	p, err := filepath.Abs(opcodesPath)
	if err != nil {
		return err
	}

	all := &Data{
		Opcodes: []Opcode{},
	}

	for _, ext := range extensions {

		log.Printf("processing %s", ext)

		o := &Data{
			Opcodes: []Opcode{},
		}
		for _, r := range []string{"rv_", "rv32_"} {
			err = fileProcess(o, filepath.Join(p, r+ext))
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					log.Printf("skipping %s%s, doesn't exist", r, ext)
				} else {
					return err
				}
			}
		}

		err = writeFile(tmplEncode, o, fmt.Sprintf("pkg/assemble/opcodes_%s.go", ext))
		if err != nil {
			return err
		}

		err = writeFile(tmplDecode, o, fmt.Sprintf("pkg/disassemble/opcodes_%s.go", ext))
		if err != nil {
			return err
		}

		err = writeFile(tmplDefs, o, fmt.Sprintf("pkg/opcodes/opcodes_%s.go", ext))
		if err != nil {
			return err
		}

		all.Opcodes = append(all.Opcodes, o.Opcodes...)
	}

	err = writeFile(tmplAssemble, all, "pkg/assemble/opcodes.go")
	if err != nil {
		return err
	}

	err = writeFile(tmplMatch, all, "pkg/disassemble/match.go")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "genopcodes",
		Usage: "generate opcodes",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "opcodes",
				Usage: "path to opcodes dir",
			},
			&cli.StringFlag{
				Name:        "extensions",
				Usage:       "extensions to generate for",
				DefaultText: DEFAULT_EXTENSIONS,
			},
		},
		Action: func(c *cli.Context) error {
			s := c.String("extensions")
			if s == "" {
				s = DEFAULT_EXTENSIONS
			}
			ext := strings.Split(s, ",")
			return process(c.String("opcodes"), ext)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
