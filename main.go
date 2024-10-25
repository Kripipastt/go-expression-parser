package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Key(m map[string]float64) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func req(stack map[string]string, values map[string]float64, lastIndex string) float64 {
	var expressionReplace = stack[lastIndex]
	for _, value := range "+-/*" {
		expressionReplace = strings.Replace(expressionReplace, string(value), "$"+string(value)+"$", -1)
	}
	var expressionSplit = strings.Split(expressionReplace, "$")
	var ind1, ind2 = expressionSplit[0], expressionSplit[2]
	var val1, val2 float64
	if strings.Contains(ind1, "@") {
		if !slices.Contains(Key(values), ind1) {
			values[ind1] = req(stack, values, ind1)
		}
		val1 = values[ind1]
	} else {
		val1, _ = strconv.ParseFloat(ind1, 64)
	}
	if strings.Contains(ind2, "@") {
		if !slices.Contains(Key(values), ind2) {
			values[ind2] = req(stack, values, ind2)
		}
		val2 = values[ind2]
	} else {
		val2, _ = strconv.ParseFloat(ind2, 64)
	}
	switch expressionSplit[1] {
	case "+":
		return val1 + val2
	case "-":
		return val1 - val2
	case "*":
		return val1 * val2
	case "/":
		return val1 / val2
	default:
		panic("error")
	}
}

func C(expression string) (float64, error) {
	var stack = map[string]string{}
	var n = 1
	//var isPriority = false
	var newStr = expression
	for _, operation := range "+-/*" {
		newStr = strings.Replace(newStr, string(operation), "@"+string(operation)+"@", -1)
	}
	var newStrArray = strings.Split(newStr, "@")
	fmt.Println(newStrArray)
	for {
		var indexArr []int
		for key, value := range newStrArray {
			if value == "*" || value == "/" {
				indexArr = append(indexArr, key)
			}
		}
		var newInd = ""
		var newArr []string
		for key, value := range newStrArray {
			if slices.Contains(indexArr, key) {
				newInd = "@" + strconv.Itoa(n)
				newArr = append(newArr, newInd)
				stack[newInd] = newStrArray[key-1] + value + newStrArray[key+1]
				newStrArray = append(newArr, newStrArray[key+2:]...)
				newArr = []string{}
				n += 1
				break
			} else if slices.Contains(indexArr, key-1) || slices.Contains(indexArr, key+1) {

			} else {
				newArr = append(newArr, value)
			}
		}
		if !slices.Contains(newStrArray, "/") && !slices.Contains(newStrArray, "*") {
			break
		}
	}
	// --------------
	for {
		var indexArr []int
		for key, value := range newStrArray {
			if value == "-" || value == "+" {
				indexArr = append(indexArr, key)
			}
		}
		var newInd = ""
		var newArr []string
		for key, value := range newStrArray {
			if slices.Contains(indexArr, key) {
				newInd = "@" + strconv.Itoa(n)
				newArr = append(newArr, newInd)
				stack[newInd] = newStrArray[key-1] + value + newStrArray[key+1]
				newStrArray = append(newArr, newStrArray[key+2:]...)
				newArr = []string{}
				n += 1
				break
			} else if slices.Contains(indexArr, key-1) || slices.Contains(indexArr, key+1) {

			} else {
				newArr = append(newArr, value)
			}
		}
		if !slices.Contains(newStrArray, "-") && !slices.Contains(newStrArray, "+") {
			break
		}
	}
	fmt.Println(newStrArray)
	fmt.Println(stack)
	result := req(stack, map[string]float64{}, newStrArray[0])
	return result, nil
}

func main() {
	fmt.Println(C("2+3"))
}
