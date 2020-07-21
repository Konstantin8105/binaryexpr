package binaryexpr

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/Konstantin8105/errors"
)

// Test binary expression in file with filename and return error if used
// not acceptable binary expression `>=` or `>`.
func Test(filename string) error {
	// positions are relative to fset
	fset := token.NewFileSet()

	// parse file
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// create new tree error
	et := errors.New(filename)

	// inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		if bin, ok := n.(*ast.BinaryExpr); ok {
			switch bin.Op {
			case token.GEQ, token.GTR: // >= or >
				et.Add(fmt.Errorf("%s:\tnot acceprable operator: %v\n",
					fset.Position(n.Pos()), bin.Op))
			}
		}
		return true
	})

	// error handling
	if et.IsError() {
		return et
	}
	return nil
}
