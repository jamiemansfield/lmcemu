package main

import (
	"github.com/urfave/cli"
	"os"
	"log"
	"net/url"
	"fmt"
	"github.com/jamiemansfield/lmcemu/asm"
	"github.com/mileusna/conditional"
	"encoding/json"
	"github.com/jamiemansfield/lmcemu/emu"
)

var (
	inputFlag = cli.StringFlag{
		Name: "input, i",
		Usage: "the input file or URI",
	}
)

type lmcFunc func(instructions []*asm.Instruction) error

func createAction(lmcFunc lmcFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		input := c.String("input")
		uri, err := url.Parse(input)
		if err != nil {
			return err
		}

		var instructions []*asm.Instruction
		if uri.Scheme == "file" || uri.Scheme == "" {
			file, err := os.Open(conditional.String(uri.Scheme == "", uri.Path, uri.Opaque))
			if err != nil {
				return err
			}
			defer file.Close()

			// Create a Parser and parse the file
			parser := asm.CreateParser()
			parser.ReadFromFile(file)

			// Assemble the program.
			insts, err := parser.Assemble()
			if err != nil {
				return err
			}
			instructions = insts
		} else
		if uri.Scheme == "builtin" {
			insts := BuiltinRegistry[uri.Opaque]
			if insts == nil {
				return fmt.Errorf("compile: a builtin program of name '%s' does not exist", uri.Opaque)
			}
			instructions = insts
		} else {
			return fmt.Errorf("compile: unsupported URI scheme ('%s') used", uri.Scheme)
		}

		return lmcFunc(instructions)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "lmcemu"
	app.Usage = "a Little Man Computer emulator"
	app.Version = "1.0.0-indev"

	app.Commands = []cli.Command{
		{
			Name:    "compile",
			Aliases: []string{"c"},
			Usage:   "compiles the given assembly file to machine code",
			Flags:   []cli.Flag {
				inputFlag,
			},
			Action:  createAction(func (instructions []*asm.Instruction) error {
				// Assemble the program into machine code
				prog, err := asm.AssembleProgram(instructions)
				if err != nil {
					return err
				}

				// Print the machine code
				fmt.Printf("%v", prog)

				return nil
			}),
		},
		{
			Name:    "dump",
			Aliases: []string{"d"},
			Usage:   "dumps the given assembly file as a JSON output",
			Flags:   []cli.Flag {
				inputFlag,
			},
			Action:  createAction(func (instructions []*asm.Instruction) error {
				// Dump in JSON form
				bytes, err := json.MarshalIndent(instructions, "", "\t")
				if err != nil {
					return err
				}
				fmt.Println(string(bytes))

				return nil
			}),
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "runs the given assembly file",
			Flags:   []cli.Flag {
				inputFlag,
			},
			Action:  createAction(func (instructions []*asm.Instruction) error {
				// Assemble the program into machine code
				prog, err := asm.AssembleProgram(instructions)
				if err != nil {
					return err
				}

				// Run
				cpu := emu.CreateLmcCpu()
				memory := emu.CreateMemory(prog)
				err = cpu.Execute(memory)
				if err != nil {
					return err
				}

				return nil
			}),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
