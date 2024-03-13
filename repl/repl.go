package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		_lexer := lexer.New(line)

		for _token := _lexer.NextToken(); _token.Type != token.EOF; _token = _lexer.NextToken() {
			fmt.Printf("%+v\n", _token)
		}
	}
}
