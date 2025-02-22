package storage

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Expression struct {
	Id            string
	Expression    string
	Status        string
	FinalTask     string
	Result        float64
	Stack         map[string]string
	CountedValues map[string]float64
}

const (
	CREATE  = "create"
	PENDING = "pending"
	FINISH  = "finish"
	REJECT  = "reject"
)

type Storage struct {
	Expressions []*Expression
}

func NewStorage() *Storage {
	return &Storage{Expressions: make([]*Expression, 0)}
}

func NewExpression(id, expression, finalTask string, stack map[string]string) *Expression {
	return &Expression{Id: id, Expression: expression, FinalTask: finalTask, Stack: stack, Status: CREATE, CountedValues: make(map[string]float64)}
}

func (s *Storage) Add(expression string, finalTask string, stack map[string]string) string {
	id := uuid.NewString()
	newExpression := NewExpression(id, expression, finalTask, stack)
	s.Expressions = append(s.Expressions, newExpression)

	return id
}

func (s *Storage) GetAll() []*Expression {
	return s.Expressions
}

func (s *Storage) GetById(id string) (*Expression, error) {
	for _, expression := range s.Expressions {
		if expression.Id == id {
			return expression, nil
		}
	}
	return &Expression{}, errors.New(fmt.Sprintf("Expression with id %s not found", id))
}

func (s *Storage) GetUnsolvedExpressions() []*Expression {
	unsolvedExpressions := make([]*Expression, 0)
	for i, expression := range s.Expressions {
		if expression.Status == CREATE || expression.Status == PENDING {
			unsolvedExpressions = append(unsolvedExpressions, s.Expressions[i])
		}
	}
	return unsolvedExpressions
}

func (s *Storage) SetStatus(expression *Expression, newStatus string) {
	expression.Status = newStatus
}

func (s *Storage) AddTaskResult(expression *Expression, taskId string, result float64) {
	expression.CountedValues[taskId] = result
}

func (s *Storage) FinishExpression(expression *Expression, result float64) {
	expression.Status = FINISH
	expression.Result = result
}

var ExStorage = NewStorage()
