package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Kripipastt/go-expression-parser/agent/config"
	"net/http"
)

type Task struct {
	Id            string  `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime int     `json:"operation_time"`
}

type TaskResponse struct {
	Task Task `json:"task"`
}

type TaskRequest struct {
	Id     string  `json:"id"`
	Result float64 `json:"result"`
}

func GetTask() (Task, error) {
	response, err := http.Get(config.Service.OrchestratorUrl + "/internal/task")
	if err != nil || response.StatusCode != 200 {
		return Task{}, errors.New("no waiting tasks")
	}
	defer response.Body.Close()

	var taskResponse TaskResponse
	decoder := json.NewDecoder(response.Body)
	decoder.Decode(&taskResponse)
	//fmt.Println(taskResponse)
	return taskResponse.Task, nil
}

func PostTask(taskId string, result float64) {
	requestBody, _ := json.Marshal(TaskRequest{Id: taskId, Result: result})
	http.Post(config.Service.OrchestratorUrl+"/internal/task", "application/json", bytes.NewReader(requestBody))
}
