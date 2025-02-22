package orchestrator

import (
	"errors"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/application/messages"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/config"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/storage"
	"github.com/Kripipastt/go-expression-parser/orchestrator/pkg/parser"
	"log"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type TaskManager struct {
	activeTask []string
	mu         sync.Mutex
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func SelectOperationTime(operation string) int {
	switch operation {
	case "+":
		return config.Service.TimeAdditionMs
	case "-":
		return config.Service.TimeSubtractionMs
	case "*":
		return config.Service.TimeMultiplicationMs
	case "/":
		return config.Service.TimeDivisionMs
	case "^":
		return config.Service.TimeExponentiationMs
	}
	return 0
}

func (tm *TaskManager) GetTask() (messages.Task, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	//fmt.Println(tm.activeTask)
	//fmt.Println(storage.ExStorage.Expressions[0])

	unsolvedExpressions := storage.ExStorage.GetUnsolvedExpressions()

	for _, currentExpression := range unsolvedExpressions {

		for taskId, strTask := range currentExpression.Stack {
			if _, ok := currentExpression.CountedValues[taskId]; !ok && !slices.Contains(tm.activeTask, currentExpression.Id+"_"+taskId) {
				var goodTask = true
				var param1, param2 float64
				var exists bool

				splitTask := parser.SplitTask(strTask)
				arg1, operation, arg2 := splitTask[0], splitTask[1], splitTask[2]
				if operation == "+" && arg2[0] == '-' {
					arg2 = arg2[1:]
					operation = "-"
				}

				if strings.Contains(arg1, "@") {
					param1, exists = currentExpression.CountedValues[strings.Replace(arg1, "-", "", -1)]
					if exists && arg1[0] == '-' {
						param1 = -param1
					}

					if !exists {
						goodTask = false
					}
				} else {
					param1, _ = strconv.ParseFloat(arg1, 64)
				}

				if strings.Contains(arg2, "@") {
					param2, exists = currentExpression.CountedValues[strings.Replace(arg2, "-", "", -1)]
					if exists && arg2[0] == '-' {
						param2 = -param2
					}

					if !exists {
						goodTask = false
					}
				} else {
					param2, _ = strconv.ParseFloat(arg2, 64)
				}

				if goodTask && (operation == "/" && param2 == 0) || (operation == "^" && param1 < 0 && !parser.CheckCorrectExponentiation(param2)) {
					log.Printf("Expression: %s reject for %f %s %f", currentExpression.Expression, param1, operation, param2)
					storage.ExStorage.SetStatus(currentExpression, storage.REJECT)
					break
				}

				if goodTask {
					if currentExpression.Status == storage.CREATE {
						storage.ExStorage.SetStatus(currentExpression, storage.PENDING)
					}
					tm.activeTask = append(tm.activeTask, currentExpression.Id+"_"+taskId)
					return messages.Task{Id: currentExpression.Id + "_" + taskId, Arg1: param1, Arg2: param2, Operation: operation, OperationTime: SelectOperationTime(operation)}, nil
				}
			}
		}
	}
	return messages.Task{}, errors.New("no waiting tasks")
}

func (tm *TaskManager) PostTaskResult(expression *storage.Expression, fullTaskId, taskId string, result float64) error {

	tm.mu.Lock()
	defer tm.mu.Unlock()
	var isActiveTask = false
	for i, task := range tm.activeTask {
		if task == fullTaskId {
			tm.activeTask = append(tm.activeTask[:i], tm.activeTask[i+1:]...)
			isActiveTask = true
			break
		}
	}
	if !isActiveTask {
		return errors.New("task with id " + taskId + " was not registered as active")
	}

	storage.ExStorage.AddTaskResult(expression, taskId, result)
	if taskId == expression.FinalTask {
		storage.ExStorage.FinishExpression(expression, result)
	}
	return nil
}

var TM = NewTaskManager()
