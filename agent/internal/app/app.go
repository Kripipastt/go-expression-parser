package app

import (
	"github.com/Kripipastt/go-expression-parser/agent/config"
	"github.com/Kripipastt/go-expression-parser/agent/internal/agent"
	"sync"
)

func RunAgents() {
	agentCount := config.Service.ComputingPower
	var wg sync.WaitGroup
	for i := 1; i <= agentCount; i++ {
		go agent.CreateAgent(i)
		wg.Add(1)
	}
	wg.Wait()
}
