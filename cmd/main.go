package main

import (
	"os"
	"log"
	"fmt"
	"alexdenkk/bf/internal/bf"
)

const HELP_MSG = `
Brainfuck compiler | By alexdenkk

Usage:
build <filename> - compile file

Commands:
, | input cell
. | print cell
> | next cell
< | previous cell
+ | plus 1 to cell value
- | minus to cell value
[ | open cycle
] | close cycle

*Cycles must be closed

Basic expressions:
,>,[<+>-]<.                                | summation
>,[<->-]<.                                 | subtraction
,>,[<[->>+>+<<<]>>>[-<<<+>>>]<<<>-]<[-]>>. | multiplication
`

func main() {
	compiler := bf.New()

	if len(os.Args) > 2 {
		err := compiler.CompileFile(os.Args[2])

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(HELP_MSG)
	}
}
