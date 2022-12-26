package main

import (
	"bytes"
	"errors"
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

const DEFAULT_EXTENSIONS = "i,zicsr"

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

		op := Opcode{}
		fields := make(map[string]int)

		for i, f := range strings.Fields(ln) {
			if i == 0 {
				op.Name = strings.Replace(f, ".", "_", -1)
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
					op.Fields = append(op.Fields, [2]string{f, fieldname})
					fields[fieldname] = 1
					found = true
				}

				if !found {
					log.Printf("!! couldn't find field for %s\n", f)
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

func process(opcodesPath string, extensions []string) error {

	msblsb = regexp.MustCompile(`\s*(?P<msb>\d+.?)\.\.(?P<lsb>\d+.?)\s*=\s*(?P<val>\d[\w]*)[\s$]*`)
	args = regexp.MustCompile(`\s?(?:(?:((?:j|b|z)?imm(?:12|20)?)(?:hi|lo)?)+|(r(?:s\d|d))+|(fm|pred|succ|csr|shamtw)+)\s?`)

	tmpl := template.Must(template.ParseFiles("hack/assemble/opcode_parser.tmpl"))

	p, err := filepath.Abs(opcodesPath)
	if err != nil {
		return err
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

		var out bytes.Buffer
		err = tmpl.Execute(&out, o)
		if err != nil {
			return err
		}

		formatted, err := format.Source(out.Bytes())
		if err != nil {
			return fmt.Errorf("failed to format: %w", err)
		}

		f, err := os.Create(fmt.Sprintf("pkg/assemble/opcodes_%s.go", ext))
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
