package agent

import (
	"github.com/Kripipastt/go-expression-parser/agent/pkg/calc"
	"github.com/Kripipastt/go-expression-parser/agent/pkg/http"
	"log"
	"time"
)

func CreateAgent(n int) {
	log.Println("Creating agent" + string(rune(n)))
	for {
		task, err := http.GetTask()
		if err != nil {
			log.Printf("Error getting task: %s", err)
			time.Sleep(5 * time.Second)
		} else {
			select {
			case <-time.After(time.Duration(task.OperationTime) * time.Millisecond):
				result := calc.CalcTask(task.Arg1, task.Arg2, task.Operation)
				http.PostTask(task.Id, result)
				log.Printf("Task with id %s (%f %s %f) finished successfully by agent %d", task.Id, task.Arg1, task.Operation, task.Arg2, n)
			}
		}
	}
}
