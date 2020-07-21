package binaryexpr

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func Test(t *testing.T, filename string) {
	// positions are relative to fset
	fset := token.NewFileSet()

	// parse file
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		t.Error(err)
		return
	}

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		if bin, ok := n.(*ast.BinaryExpr); ok {
			switch bin.Op {
			case token.GEQ, token.GTR: // >= or >
				t.Errorf("%s:\tOp: %v\n", fset.Position(n.Pos()), bin.Op)
			}
		}
		return true
	})
}
