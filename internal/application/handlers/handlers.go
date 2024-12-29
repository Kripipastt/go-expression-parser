package handlers

import (
	"encoding/json"
	"github.com/Kripipastt/go-expression-parser/pkg/parser"
	"log"
	"net/http"
)

type Request struct {
	Expression string `json:"expression" example:"2 + 2 * 2"`
}

type Response struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// CalcHandler
//
//	@Summary      Calculate expression
//	@Description  Parse and calculate expression
//	@Tags         expression
//	@Accept       json
//	@Produce      json
//	@Param        expression    body     Request  true  "Expression for parse and calc"
//	@Success      200  {object}  Response
//	@Failure      422  {object}  ErrorResponse
//	@Failure      500  {object}  ErrorResponse
//	@Router       /api/v1/calculate [post]
func CalcHandler(w http.ResponseWriter, r *http.Request) {
	var request = Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Internal server error"})
		log.Println("ERROR: ", err)
		return
	}

	defer r.Body.Close()

	result, err := parser.Calc(request.Expression)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Expression is not valid"})
		log.Println("ERROR: ", err)
		return
	}

	json.NewEncoder(w).Encode(Response{Result: result})
}

// PingHandler
// @Summary Ping
// @Description Ping for healthcheck
// @Tags Other
// @Produce json
// @Success 	200 {string} string
// @Router		/ping [get]
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
