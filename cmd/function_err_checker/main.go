package main

import (
	function_err_checker "github.com/k3forx/function-err-checker"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(function_err_checker.Analyzer)
}
