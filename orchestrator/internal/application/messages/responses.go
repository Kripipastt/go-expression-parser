package messages

type ResponseExpressionId struct {
	Id string `json:"id"`
}

type ResponseExpression struct {
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

type ResponseAllExpression struct {
	Expressions []ResponseExpression `json:"expressions"`
}

type ResponseOneExpression struct {
	Expression ResponseExpression `json:"expression"`
}

type ResponseError struct {
	Error string `json:"error"`
}

type Task struct {
	Id            string  `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime int     `json:"operation_time"`
}

type ResponseTask struct {
	Task Task `json:"task"`
}

type ResponseOk struct {
	Status string
}
