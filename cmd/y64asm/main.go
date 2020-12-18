package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Joakker/y64asm"
)

var file io.Reader

func init() {
	silent := flag.Bool("s", false, "Don't output to the terminal")
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
	lexer := y64asm.NewLexer(file)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == y64asm.EOF {
			break
		}
		fmt.Printf("%d:%d\t%s\t%s\n", pos.Row, pos.Col, tok, lit)
	}
}
