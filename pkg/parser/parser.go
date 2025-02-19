package parser

import (
	maps2 "golang.org/x/exp/maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

const (
	ALLOWEDCHAR       = "0123456789+-*/^()."
	OPERATIONS        = "+/*^"
	EXTENDEDOPERATION = "+/*^()"
)

func deleteSpecificElement[T comparable](array []T, deletedElement T) []T {
	var newArray []T
	for _, el := range array {
		if el != deletedElement {
			newArray = append(newArray, el)
		}
	}
	return newArray
}

func findIndexesOfElements[T comparable](array []T, needsElements []T) []int {
	var indexes []int
	for i, el := range array {
		if slices.Contains(needsElements, el) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func recursiveStackCounting(stack map[string]string, countedValues map[string]float64, lastIndex string) (float64, error) {
	var expressionReplace = stack[lastIndex]
	for _, value := range OPERATIONS {
		expressionReplace = strings.Replace(expressionReplace, string(value), "$"+string(value)+"$", -1)
	}
	var expressionSplit = strings.Split(expressionReplace, "$")

	if len(expressionSplit) != 3 {
		return 0, unknownOperator
	}

	var indexes = [2]string{expressionSplit[0], expressionSplit[2]}
	var values = [2]float64{}
	var err error = nil
	for i, index := range indexes {
		if strings.Contains(index, "@") {
			if !slices.Contains(maps2.Keys(countedValues), index) {
				if string(index[0]) == "-" {
					countedValues[index], err = recursiveStackCounting(stack, countedValues, index[1:])
					countedValues[index] = -countedValues[index]
				} else {
					countedValues[index], err = recursiveStackCounting(stack, countedValues, index)
				}
			}
			values[i] = countedValues[index]
		} else {
			values[i], _ = strconv.ParseFloat(index, 64)
		}
	}
	switch expressionSplit[1] {
	case "+":
		return values[0] + values[1], err
	case "*":
		return values[0] * values[1], err
	case "/":

		if values[1] == 0 {
			return 0, divideByZero
		}
		return values[0] / values[1], err
	case "^":
		return math.Pow(values[0], values[1]), err
	default:
		return 0, unknownOperator
	}
}

func checkCorrectExpression(expression string) error {
	if len(expression) == 0 {
		return emptyExpression
	}
	if (strings.Contains(OPERATIONS, string(expression[0])) && !strings.Contains("-", string(expression[0]))) || strings.Contains(OPERATIONS, string(expression[len(expression)-1])) {
		return incorrectExpression
	}
	var previousElementIsOperation = false
	for _, el := range expression {
		if !strings.Contains(ALLOWEDCHAR, string(el)) {
			return unallowableChar
		}
		if strings.Contains(OPERATIONS, string(el)) {
			if previousElementIsOperation {
				return severalOperations
			}
			previousElementIsOperation = true
		} else {
			previousElementIsOperation = false
		}
	}
	return nil
}

func splitStringIntoTokens(str, tokens string) []string {
	for _, operation := range tokens {
		str = strings.Replace(str, string(operation), "#"+string(operation)+"#", -1)
		str = strings.Replace(str, "##", "#", -1)
	}
	return strings.Split(str, "#")
}

func Calc(expression string) (float64, error) {
	// Отчищаем выражения от лишних пробнлов
	expression = strings.Join(deleteSpecificElement(strings.Split(expression, ""), " "), "")

	// Проверяем на корректность
	err := checkCorrectExpression(expression)
	if err != nil {
		return 0, err
	}

	// Заменяет все - на +- для более удобных вычислений
	expression = strings.Replace(expression, "--", "", -1)
	expression = strings.Replace(expression, "-", "+-", -1)
	expression = strings.Replace(expression, "(+-", "(-", -1)
	if string(expression[0]) == "+" {
		expression = expression[1:]
	}

	var stack = map[string]string{}
	var tokens = deleteSpecificElement(splitStringIntoTokens(expression, EXTENDEDOPERATION), "")

	for _, token := range tokens {
		if strings.Count(token, ".") > 1 {
			return 0, incorrectFractional
		}
	}

	//fmt.Println(tokens)
	for {
		//fmt.Println(tokens)
		for _, operations := range [][]string{{"^"}, {"*", "/"}, {"+"}} {
			for {
				var maxParenthesis, currentParenthesis = 0, 0
				for _, token := range tokens {
					if token == "(" {
						currentParenthesis += 1
					} else if token == ")" {
						currentParenthesis -= 1
					}
					if currentParenthesis < 0 {
						return 0, incorrectParenthesis
					}
					maxParenthesis = max(maxParenthesis, currentParenthesis)
				}
				if currentParenthesis != 0 {
					return 0, incorrectParenthesis
				}
				tokens = deleteSpecificElement(tokens, "")
				var indexesOperations = findIndexesOfElements(tokens, operations)
				var unusedTokens []string
				var unfilteredTokens []string
				var isWrite = 2
				currentParenthesis = 0
				for key, value := range tokens {
					if isWrite == 2 {
						if slices.Contains(indexesOperations, key) && !strings.Contains("()", tokens[key-1]) && !strings.Contains("()", tokens[key+1]) && currentParenthesis == maxParenthesis {
							var newInd = "@" + strconv.Itoa(len(stack))
							unusedTokens = append(unusedTokens, newInd)
							if key+2 < len(tokens) {
								unusedTokens = append(unusedTokens, tokens[key+2])
							}
							stack[newInd] = tokens[key-1] + value + tokens[key+1]
							isWrite = 0
							//unfilteredTokens = append(unusedTokens, tokens[key+2:]...)
							//break
						} else if !(slices.Contains(indexesOperations, key+1) && !strings.Contains("()", value) && currentParenthesis == maxParenthesis) {
							unusedTokens = append(unusedTokens, value)
						}
					} else {
						isWrite++
					}
					if value == "(" {
						currentParenthesis += 1
					} else if value == ")" {
						currentParenthesis -= 1
					}
				}
				//fmt.Println(operations)
				//fmt.Println(unusedTokens)

				var stringOfTokens = strings.Join(unfilteredTokens, "")
				if len(unfilteredTokens) == 0 {
					stringOfTokens = strings.Join(unusedTokens, "")
				}

				var formattedString = stringOfTokens
				for {
					var parenthesisIndexes []int
					for i, el := range stringOfTokens {
						if string(el) == "(" {
							parenthesisIndexes = append(parenthesisIndexes, i)
						} else if string(el) == ")" {
							var stringSlice = stringOfTokens[parenthesisIndexes[len(parenthesisIndexes)-1] : i+1]
							parenthesisIndexes = parenthesisIndexes[:len(parenthesisIndexes)-1]
							if len(findIndexesOfElements(strings.Split(stringSlice, ""), strings.Split(OPERATIONS, ""))) == 0 {
								formattedString = strings.Replace(stringOfTokens, stringSlice, stringSlice[1:len(stringSlice)-1], 1)
								break
							}
						}
					}
					if formattedString == stringOfTokens {
						break
					}
					stringOfTokens = formattedString
				}
				stringOfTokens = strings.Replace(stringOfTokens, "--", "", -1)
				var filteredTokens = deleteSpecificElement(splitStringIntoTokens(stringOfTokens, EXTENDEDOPERATION), "")

				if len(tokens) == len(filteredTokens) {
					break
				}
				tokens = filteredTokens
				if !slices.Equal(filteredTokens, unfilteredTokens) && slices.Equal(operations, []string{"+"}) {
					break
				}
			}
		}
		tokens = deleteSpecificElement(tokens, "")
		if len(tokens) == 1 {
			break
		}
		//fmt.Println("Stack:", stack)
	}
	//fmt.Println(tokens)
	//fmt.Println(stack)
	result, err := recursiveStackCounting(stack, map[string]float64{}, tokens[0])
	return result, err
}
