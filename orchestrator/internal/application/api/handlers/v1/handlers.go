package handlers

import (
	"encoding/json"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/application/messages"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/storage"
	"github.com/Kripipastt/go-expression-parser/orchestrator/pkg/parser"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

// CalcHandler
//
//	@Summary      Calculate expression
//	@Description  Parse and calculate expression
//	@Tags         Expression
//	@Accept       json
//	@Produce      json
//	@Param        expression    body     messages.RequestAddExpression  true  "Expression for parse and calc"
//	@Success      201  {object}  messages.ResponseExpressionId
//	@Failure      422  {object}  messages.ResponseError
//	@Failure      500  {object}  messages.ResponseError
//	@Router       /api/v1/calculate [post]
func CalcHandler(w http.ResponseWriter, r *http.Request) {
	var request = messages.RequestAddExpression{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(messages.ResponseError{Error: "Internal server error"})
		log.Println("ERROR: ", err)
		return
	}

	defer r.Body.Close()

	stack, finalTask, err := parser.Parse(request.Expression)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(messages.ResponseError{Error: err.Error()})
		log.Println("ERROR: ", err)
		return
	}

	expressionId := storage.ExStorage.Add(request.Expression, finalTask, stack)
	json.NewEncoder(w).Encode(messages.ResponseExpressionId{Id: expressionId})
	w.WriteHeader(http.StatusCreated)
}

// GetExpressionsHandler
//
//	@Summary      Get expressions
//	@Description  Get all expressions
//	@Tags         Expression
//	@Accept       json
//	@Produce      json
//	@Success      200  {object}  messages.ResponseAllExpression
//	@Failure      500  {object}  messages.ResponseError
//	@Router       /api/v1/expressions [get]
func GetExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	expression := storage.ExStorage.GetAll()
	formattedExpressions := make([]messages.ResponseExpression, 0)
	for _, exp := range expression {
		formattedExpressions = append(formattedExpressions, messages.ResponseExpression{Expression: exp.Expression, Id: exp.Id, Status: exp.Status, Result: exp.Result})
	}
	response := messages.ResponseAllExpression{
		Expressions: formattedExpressions,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(messages.ResponseError{Error: "Internal server error"})
		log.Println("ERROR: ", err)
	}
}

// GetOneExpressionHandler
//
//		@Summary      Get one expression
//		@Description  Get one expression
//		@Tags         Expression
//		@Accept       json
//		@Produce      json
//	 @Param 		  id path string true "Expression id"
//		@Success      200  {object}  messages.ResponseOneExpression
//	 @Failure  	  404  {object}  messages.ResponseError
//		@Failure      500  {object}  messages.ResponseError
//		@Router       /api/v1/expressions/{id} [get]
func GetOneExpressionHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	expressionId := chi.URLParam(r, "id")
	exp, err := storage.ExStorage.GetById(expressionId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(messages.ResponseError{Error: err.Error()})
		log.Println("ERROR: ", err)
		return
	}

	response := messages.ResponseOneExpression{Expression: messages.ResponseExpression{Expression: exp.Expression, Id: exp.Id, Status: exp.Status, Result: exp.Result}}
	err = encoder.Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(messages.ResponseError{Error: "Internal server error"})
		log.Println("ERROR: ", err)
	}
}
