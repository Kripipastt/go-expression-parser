package messages

type RequestAddExpression struct {
	Expression string `json:"expression" example:"2 + 2 * 2"`
}

type RequestPostTaskAnswer struct {
	Id     string  `json:"id"`
	Result float64 `json:"result"`
}
