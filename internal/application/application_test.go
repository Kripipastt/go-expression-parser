package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	var testCases = []struct {
		name       string
		expression string
		expected   float64
	}{
		{name: "Easy expression",
			expression: "2 + 2",
			expected:   4},
		{name: "Hard expression",
			expression: "2 * 2 / 4 + 3 - 26 + 3 - (5 - 50 * 3 / 10)",
			expected:   -9},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(Request{Expression: testCase.expression})
			request := httptest.NewRequest(http.MethodGet, "http://localhost/calc", b)
			CalcHandler(w, request)
			result := w.Result()
			defer result.Body.Close()
			data := Response{}
			json.NewDecoder(result.Body).Decode(&data)
			if data.Result != testCase.expected || result.StatusCode != http.StatusOK {
				t.Errorf("wrong result! Expected: %f, status: %d, got %f, status: %d", testCase.expected, http.StatusOK, data.Result, result.StatusCode)
			}
		})
	}
}

func TestCalcHandlerBadRequest(t *testing.T) {
	var testCases = []struct {
		name       string
		expression string
	}{
		{name: "Extra characters",
			expression: "Easter egg for the examiners :3"},
		{name: "Division bu zero",
			expression: "2 - 2 / (2 - 2)"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(Request{Expression: testCase.expression})
			request := httptest.NewRequest(http.MethodGet, "http://localhost/calc", b)
			CalcHandler(w, request)
			result := w.Result()
			if result.StatusCode != http.StatusUnprocessableEntity {
				t.Errorf("wrong result! Expeted status: %d, status: %d", http.StatusUnprocessableEntity, result.StatusCode)
			}
		})
	}
}
