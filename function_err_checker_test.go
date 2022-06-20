package function_err_checker_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/k3forx/function_err_checker"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, function_err_checker.Analyzer, "a")
}
