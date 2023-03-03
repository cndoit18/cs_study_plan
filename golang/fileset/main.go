package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {

	src := []byte(`printf("Hello, World!")`)
	fset := token.NewFileSet()
	file := fset.AddFile("hello.go", fset.Base(), len(src))

	s := scanner.Scanner{}
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
