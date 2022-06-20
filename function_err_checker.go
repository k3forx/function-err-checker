package function_err_checker

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "function_err_checker is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "checker",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if ifStmt, ok := n.(*ast.IfStmt); ok {
			if assignStmt, ok := ifStmt.Init.(*ast.AssignStmt); ok {
				if callExpr, ok := assignStmt.Rhs[0].(*ast.CallExpr); ok {
					if ident, ok := callExpr.Fun.(*ast.Ident); ok {
						if ident.Name == "Translate" {
							if ifSt, ok := ifStmt.Body.List[0].(*ast.IfStmt); ok {
								if ce, ok := ifSt.Cond.(*ast.CallExpr); ok {
									if ide, ok := ce.Fun.(*ast.Ident); ok {
										if ide.Name == "IsTypeError" {
											pass.Reportf(ide.NamePos, "checked!")
										}
									}
								}
							}
						}
					}
				}
			}
		}
		// switch n := n.(type) {
		// case *ast.Ident:
		// 	if n.Name == "gopher" {
		// 		pass.Reportf(n.Pos(), "identifier is gopher")
		// 	}
		// }
	})

	return nil, nil
}
