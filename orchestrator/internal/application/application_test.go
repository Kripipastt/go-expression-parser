package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/application/messages"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExecuteRequest(req *http.Request, s *Application) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func CheckResponse(t *testing.T, expectedCode, actualCode int, expected, actual interface{}) {
	if expectedCode != actualCode || expected != actual {
		t.Errorf("Expected code %d, got %d. Expected body %f, got %f", actualCode, expectedCode, expected, actual)
	}
}

func TestCalcHandler(t *testing.T) {
	var testCases = []struct {
		name       string
		expression string
		expected   float64
	}{
		//{name: "Easy expression",
		//	expression: "2 + 2",
		//	expected:   4},
		//{name: "Hard expression",
		//	expression: "2 * 2 / 4 + 3 - 26 + 3 - (5 - 50 * 3 / 10)",
		//	expected:   -9},
	}

	app := NewApplication()
	app.MountHandlers()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(messages.RequestAddExpression{Expression: testCase.expression})
			request := httptest.NewRequest(http.MethodPost, "http://localhost/api/v1/calculate", b)

			response := ExecuteRequest(request, app)
			data := messages.ResponseExpressionId{}
			json.NewDecoder(response.Body).Decode(&data)
			fmt.Println(data)
			//CheckResponse(t, http.StatusOK, response.Code, testCase.expected, data.Result)
		})
	}
}

func TestCalcHandlerBadRequest(t *testing.T) {
	var testCases = []struct {
		name       string
		expression string
	}{
		//{name: "Extra characters",
		//	expression: "Easter egg for the examiners :3"},
		//{name: "Division bu zero",
		//	expression: "2 - 2 / (2 - 2)"},
	}

	app := NewApplication()
	app.MountHandlers()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(messages.RequestAddExpression{Expression: testCase.expression})
			request := httptest.NewRequest(http.MethodPost, "http://localhost/api/v1/calculate", b)

			response := ExecuteRequest(request, app)
			data := messages.ResponseExpressionId{}
			json.NewDecoder(response.Body).Decode(&data)
			CheckResponse(t, http.StatusUnprocessableEntity, response.Code, 0, 0)
		})
	}
}
