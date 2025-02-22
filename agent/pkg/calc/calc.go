package calc

import (
	"log"
	"math"
)

func CalcTask(arg1, arg2 float64, operation string) float64 {
	switch operation {
	case "+":
		return arg1 + arg2
	case "-":
		return arg1 - arg2
	case "*":
		return arg1 * arg2
	case "/":

		return arg1 / arg2
	case "^":
		return math.Pow(arg1, arg2)
	default:
		log.Fatalf("Unknown operation: %f %s %f", arg1, operation, arg2)
		return 0
	}
}
