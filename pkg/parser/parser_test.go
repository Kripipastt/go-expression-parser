package parser

import (
	"errors"
	"testing"
)

func TestCalc(t *testing.T) {
	var cases = []struct {
		name, expression string
		answer           float64
		err              error
	}{{"Easy expression", "2+2-2", 2, nil},
		{"Expression priority", "2+2*2", 6, nil},
		{"Hard expression", "4+3/2+4*2-5/5", 12.50, nil},
		{"Parenthesis work", "(2+2)*2", 8, nil},
		{"Parenthesis work 2", "(2 + 3) * 1 + 1", 6, nil},
		{"Hard expression with parenthesis", "(3+(2+3)*1+(2+2))/1+(4-3)", 13, nil},
		{"Check space delete", "3   +(2      + 1 )  ", 6, nil},
		{"Check division by zero error", "(4-2)/(10*0)", 0, divideByZero},
		{"Incorrect parenthesis", "(4+3", 0, incorrectParenthesis},
		{"Unallowable character", "4+3-4x+2/1", 0, unallowableChar},
		{"Many parenthesis", "(((4))+3)", 7, nil},
		{"Empty expression", "", 0, emptyExpression},
		{"Incorrect expression", "*2+3", 0, incorrectExpression},
		{"Check pow", "(2 +2)^2", 16, nil},
		{"Check unary minus sign", "- 2 + 3 - (-4)", 5, nil},
		{"Stupid parenthesis", "-((2 + (2))) + 2 + 2", 0, nil},
		{"Fractional numbers", "(0.5 + 9^0.5) / 3.5", 1, nil},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := Calc(testCase.expression)
			if result != testCase.answer || !errors.Is(err, testCase.err) {
				t.Errorf("Calc(%s) = %f, %s; want %f, %s", testCase.expression, result, err, testCase.answer, testCase.err)
			}
		})
	}
}
