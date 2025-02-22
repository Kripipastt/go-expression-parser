package handlers

import (
	"encoding/json"
	"errors"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/application/messages"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/orchestrator"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/storage"
	"log"
	"net/http"
	"strings"
)

// GetTaskHandler
//
//			@Summary      Get task
//			@Description  Get task
//			@Tags         Internal
//			@Accept       json
//			@Produce      json
//			@Success      200  {object}  messages.ResponseTask
//	 	    @Failure  	  404  {object}  messages.ResponseError
//			@Failure      500  {object}  messages.ResponseError
//			@Router       /internal/task [get]
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	task, err := orchestrator.TM.GetTask()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(messages.ResponseError{Error: err.Error()})
		return
	}
	encoder.Encode(messages.ResponseTask{Task: task})
}

// PostTaskResultHandler
//
//		@Summary      Post task result
//		@Description  Post task result
//		@Tags         Internal
//		@Accept       json
//		@Produce      json
//		@Param        task    body     messages.RequestPostTaskAnswer  true  "Answer for task"
//		@Success      200
//	    @Failure  	  404  {object}  messages.ResponseError
//		@Failure      500  {object}  messages.ResponseError
//		@Router       /internal/task [post]
func PostTaskResultHandler(w http.ResponseWriter, r *http.Request) {
	request := messages.RequestPostTaskAnswer{}
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)
	defer r.Body.Close()

	splitId := strings.Split(request.Id, "_")
	if len(splitId) != 2 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err := errors.New("incorrect id " + request.Id)
		encoder.Encode(messages.ResponseError{Error: err.Error()})
		log.Println("ERROR: ", err)
		return
	}
	expressionId, taskId := splitId[0], splitId[1]
	currentExpression, err := storage.ExStorage.GetById(expressionId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(messages.ResponseError{Error: err.Error()})
		log.Println("ERROR: ", err)
		return
	}

	err = orchestrator.TM.PostTaskResult(currentExpression, request.Id, taskId, request.Result)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		encoder.Encode(messages.ResponseError{Error: err.Error()})
		log.Println("ERROR: ", err)
		return
	}
	encoder.Encode(messages.ResponseOk{Status: "ok"})
}
