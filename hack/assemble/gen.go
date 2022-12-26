package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
	"golang.org/x/exp/maps"
)

type Opcode struct {
	Name      string
	Args      []string
	Fields    [][2]string
	Bitfields [][2]string
}

type Data struct {
	Opcodes []Opcode
}

var msblsb, args *regexp.Regexp

func fileProcess(data *Data, filename string) error {

	log.Print(filename)

	d, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(d), "\n")

	for _, l := range lines {
		ln := strings.TrimSpace(l)
		ln, _, _ = strings.Cut(ln, "#")

		// don't use blanks or care about psuedo-ops rn
		if ln == "" || strings.HasPrefix(ln, "$pseudo_op") {
			continue
		}

		op := Opcode{}
		fields := make(map[string]int)

		for i, f := range strings.Fields(ln) {
			fmt.Println(f)
			if i == 0 {
				op.Name = f
				continue
			}

			subs := msblsb.FindStringSubmatch(f)
			if subs == nil {
				for n, fieldname := range args.FindStringSubmatch(f) {
					if fieldname == "" || n == 0 {
						continue
					}

					fieldname = strings.ToUpper(fieldname)

					op.Fields = append(op.Fields, [2]string{f, fieldname})
					fields[fieldname] = 1
				}
			} else {
				op.Bitfields = append(op.Bitfields, [2]string{
					subs[2], subs[3],
				})
			}
		}

		op.Args = maps.Keys(fields)
		data.Opcodes = append(data.Opcodes, op)

	}

	return nil

}

func process(opcodesPath string) error {

	msblsb = regexp.MustCompile(`\s*(?P<msb>\d+.?)\.\.(?P<lsb>\d+.?)\s*=\s*(?P<val>\d[\w]*)[\s$]*`)
	args = regexp.MustCompile(`\s?(?:(?:(?:j|b)?(imm(?:12|20))(?:hi|lo)?)+|(r(?:s\d|d))+)\s?`)

	o := &Data{
		Opcodes: []Opcode{},
	}
	p, err := filepath.Abs(opcodesPath)
	if err != nil {
		return err
	}

	files, err := os.ReadDir(p)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		n := filepath.Join(p, f.Name())

		if strings.HasPrefix(f.Name(), "rv_i") {
			err = fileProcess(o, n)
			if err != nil {
				return err
			}
		}

	}
	tmpl := template.Must(template.ParseFiles("hack/assemble/opcode_parser.tmpl"))

	var out bytes.Buffer
	err = tmpl.Execute(&out, o)
	if err != nil {
		return err
	}

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		//return fmt.Errorf("failed to format: %w", err)
		formatted = out.Bytes()
	}

	f, err := os.Create("pkg/assemble/opcodes.go")
	if err != nil {
		return err
	}

	_, err = f.WriteString(string(formatted))
	if err != nil {
		return err
	}

	err = f.Close()
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
		},
		Action: func(c *cli.Context) error {

			return process(c.String("opcodes"))
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
