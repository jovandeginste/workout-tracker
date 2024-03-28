//+build !novaluate

package units

import "github.com/Knetic/govaluate"

// NewConversion registers a new conversion formula from one Unit to another
func NewConversion(from, to Unit, formula string) {
	expr, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		panic(err)
	}

	// create conversion function
	fn := func(x float64) float64 {
		params := make(map[string]interface{})
		params["x"] = x

		res, err := expr.Evaluate(params)
		if err != nil {
			panic(err)
		}
		return res.(float64)
	}

	NewConversionFromFn(from, to, fn, formula)
}
