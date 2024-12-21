package parser

import "errors"

var (
	incorrectParenthesis = errors.New("incorrect parenthesis. The parenthesis are not closed or their order is broken")
	unknownOperator      = errors.New("unknown mathematical operator")
	divideByZero         = errors.New("divide by zero")
	unallowableChar      = errors.New("unallowable character. Check the expression")
	severalOperations    = errors.New("several operations in a row")
	emptyExpression      = errors.New("empty expression")
	incorrectExpression  = errors.New("incorrect expression")
)
