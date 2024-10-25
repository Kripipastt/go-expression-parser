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
	//fmt.Println(lastIndex)
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

func Calc(expression string) (float64, error) {
	var stack = map[string]string{}
	var n = 1
	//var isPriority = false
	var newStr = expression
	for _, operation := range "+-/*()" {
		newStr = strings.Replace(newStr, string(operation), "@"+string(operation)+"@", -1)
		newStr = strings.Replace(newStr, "@@", "@", -1)
	}
	var newStrArray = strings.Split(newStr, "@")
	var clearSpace []string
	for _, value := range newStrArray {
		if value != "" {
			clearSpace = append(clearSpace, value)
		}
	}
	newStrArray = clearSpace
	clearSpace = []string{}
	fmt.Println(newStrArray)
	for {
		fmt.Println(newStrArray)
		for {
			var maxSkob, currentSkob = 0, 0
			for _, value := range newStrArray {
				if value == "(" {
					currentSkob += 1
				} else if value == ")" {
					currentSkob -= 1
				}
				maxSkob = max(maxSkob, currentSkob)
			}
			for _, value := range newStrArray {
				if value != "" {
					clearSpace = append(clearSpace, value)
				}
			}
			newStrArray = clearSpace
			clearSpace = []string{}
			var indexArr []int
			var skobArr []int
			for key, value := range newStrArray {
				if value == "*" || value == "/" {
					indexArr = append(indexArr, key)
				} else if value == "(" || value == ")" {
					skobArr = append(skobArr, key)
				}
			}
			var newInd = ""
			var newArr []string
			var newStrArray2 []string
			clearSpace = []string{}
			currentSkob = 0
			for key, value := range newStrArray {
				if slices.Contains(indexArr, key) && !strings.Contains("()", newStrArray[key-1]) && !strings.Contains("()", newStrArray[key+1]) && currentSkob == maxSkob {
					newInd = "@" + strconv.Itoa(n)
					newArr = append(newArr, newInd)
					stack[newInd] = newStrArray[key-1] + value + newStrArray[key+1]
					newStrArray2 = append(newArr, newStrArray[key+2:]...)
					n += 1
					break
				} else if slices.Contains(indexArr, key+1) && !strings.Contains("()", value) && currentSkob == maxSkob {
				} else {
					if value == "(" {
						currentSkob += 1
					} else if value == ")" {
						currentSkob -= 1
					}
					newArr = append(newArr, value)
				}
			}
			//fmt.Println("start", newStrArray2, len(newStrArray2))
			// TODO: ((3)) не работает из-за того, что замена идёт у newStrArr2, он пустой, а вся строка в newArr
			newStrArrayR := strings.Join(newStrArray2, "#")
			fmt.Println(newStrArrayR)
			for key := range stack {
				newStrArrayR = strings.Replace(newStrArrayR, "(#"+key+"#)", "#"+key+"#", -1)
				newStrArrayR = strings.Replace(newStrArrayR, "##", "#", -1)
			}
			fmt.Println(newStrArrayR)
			newStrArray2 = strings.Split(newStrArrayR, "#")
			fmt.Println(newStrArray, newStrArray2, newArr, len(newStrArray2))
			if len(newStrArray2) == 1 && newStrArray2[0] == "" {
				newStrArray = newArr
				break
			}
			newStrArray = newStrArray2
			//fmt.Println(newStrArray, newStrArray2)
			//if len(newStrArray2) == 0 && len(newArr) != 0 {
			//	newStrArray = newArr
			//	break
			//}
			//newStrArray = newStrArray2
			//if slices.Equal(newStrArray, newStrArray2) && len(newStrArray) == 1 {
			//	fmt.Println(newStrArray, newStrArray2, newArr)
			//	newStrArray = newArr
			//	break
			//}
		}
		//break
		// --------------
		var maxSkob, currentSkob = 0, 0
		for _, value := range newStrArray {
			if value == "(" {
				currentSkob += 1
			} else if value == ")" {
				currentSkob -= 1
			}
			maxSkob = max(maxSkob, currentSkob)
		}
		for {
			clearSpace = []string{}
			for _, value := range newStrArray {
				if value != "" {
					clearSpace = append(clearSpace, value)
				}
			}
			newStrArray = clearSpace
			clearSpace = []string{}
			var indexArr []int
			var skobArr []int
			for key, value := range newStrArray {
				if value == "-" || value == "+" {
					indexArr = append(indexArr, key)
				} else if value == "(" || value == ")" {
					skobArr = append(skobArr, key)
				}
			}
			var newInd = ""
			var newArr []string
			var newStrArray2 []string
			currentSkob = 0
			for key, value := range newStrArray {
				if slices.Contains(indexArr, key) && !strings.Contains("()", newStrArray[key-1]) && !strings.Contains("()", newStrArray[key+1]) && currentSkob == maxSkob {
					newInd = "@" + strconv.Itoa(n)
					newArr = append(newArr, newInd)
					stack[newInd] = newStrArray[key-1] + value + newStrArray[key+1]
					newStrArray2 = append(newArr, newStrArray[key+2:]...)
					newArr = []string{}
					n += 1
					break
				} else if slices.Contains(indexArr, key+1) && !strings.Contains("()", value) && currentSkob == maxSkob {
				} else {
					if value == "(" {
						currentSkob += 1
					} else if value == ")" {
						currentSkob -= 1
					}
					newArr = append(newArr, value)
				}
			}
			//fmt.Println("start", newStrArray2, len(newStrArray2))
			newStrArrayR := strings.Join(newStrArray2, "#")
			fmt.Println(newStrArrayR)
			for key := range stack {
				newStrArrayR = strings.Replace(newStrArrayR, "(#"+key+"#)", "#"+key+"#", -1)
				newStrArrayR = strings.Replace(newStrArrayR, "##", "#", -1)
			}
			fmt.Println(newStrArrayR)
			newStrArray2 = strings.Split(newStrArrayR, "#")
			//fmt.Println("end", newStrArray2, len(newStrArray2))
			if len(newStrArray2) == 1 && newStrArray2[0] == "" {
				newStrArray = newArr
				break
			}
			newStrArray = newStrArray2
			//if len(newStrArray2) == 0 && len(newArr) != 0 {
			//	newStrArray = newArr
			//	break
			//}
			//newStrArray = newStrArray2
			//if slices.Equal(newStrArray, newStrArray2) && len(newStrArray) == 1 {
			//	newStrArray = newArr
			//	break
			//}
		}
		clearSpace = []string{}
		for _, value := range newStrArray {
			if value != "" {
				clearSpace = append(clearSpace, value)
			}
		}
		newStrArray = clearSpace
		clearSpace = []string{}
		if len(newStrArray) == 1 {
			break
		}
		//break
		fmt.Println("Stack:", stack)
	}
	fmt.Println(newStrArray)
	fmt.Println(stack)
	result := req(stack, map[string]float64{}, newStrArray[0])
	return result, nil
}

func main() {
	fmt.Println(Calc("(3+(2+3)*1+(2+2))/1+(4-3)"))
	fmt.Println(Calc("(3+1)*3"))
}
