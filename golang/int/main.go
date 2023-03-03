package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	lit9527 := &ast.BasicLit{
		Kind:  token.INT,
		Value: "9527",
	}
	ast.Print(nil, lit9527)

	expr, _ := parser.ParseExpr("x = 9527")
	ast.Print(nil, expr)
}
