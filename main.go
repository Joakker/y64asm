package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Joakker/y64asm/parser"
)

var file io.Reader

func init() {
	silent := flag.Bool("s", false, "Don't output to the terminal")
	run := flag.NewFlagSet("run", flag.ExitOnError)
	run.Bool("test", true, "A test flag")
	switch os.Args[1] {
	case "run":
		run.Parse(os.Args[2:])
	}
	flag.Parse()
	filename := flag.Arg(0)
	if len(filename) == 0 {
		if !*silent {
			log.Println("Taking input from stdin")
		}
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	l := parser.NewLexer(file)
	for {
		pos, tok, lit := l.Lex()
		if tok == parser.EOF {
			break
		}
		fmt.Printf("%d:%d\t%s\t%s\n", pos.Row, pos.Col, tok, lit)
	}
}
