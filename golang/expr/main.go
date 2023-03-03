package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func main() {
	expr, _ := parser.ParseExpr("1+1*2")
	fmt.Println(Expr(expr))
}

func Expr(expr ast.Expr) float64 {
	switch exp := expr.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.BasicLit:
		f, _ := strconv.ParseFloat(exp.Value, 64)
		return f
	}
	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) float64 {
	switch exp.Op {
	case token.ADD:
		return Expr(exp.X) + Expr(exp.Y)
	case token.MUL:
		return Expr(exp.X) * Expr(exp.Y)
	}
	return 0
}
