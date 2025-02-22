package errors

import "errors"

var (
	IncorrectParenthesis = errors.New("incorrect parenthesis. The parenthesis are not closed or their order is broken")
	UnknownOperator      = errors.New("unknown mathematical operator")
	DivideByZero         = errors.New("divide by zero")
	UnallowableChar      = errors.New("unallowable character. Check the expression")
	SeveralOperations    = errors.New("several operations in a row")
	EmptyExpression      = errors.New("empty expression")
	IncorrectExpression  = errors.New("incorrect expression")
	IncorrectFractional  = errors.New("incorrect fractional numbers")
)
